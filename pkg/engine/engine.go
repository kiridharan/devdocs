package engine

import (
	"bytes"
	"context"
	"fmt"
	"sort"
	"strings"

	"github.com/kiridharan/devdoc/pkg/llm"
	"github.com/kiridharan/devdoc/pkg/parser"
)

// Prompt templates
const (
	PromptDocString = `Analyze the following code snippet and generate a comprehensive docstring in the standard format for the language.
Focus on:
1. What the code does.
2. The parameters and return values.
3. Any edge cases or errors thrown.
4. Do NOT include opening/closing quotes (e.g., """ or /** */) in your response unless asked. Just the content.
   Actually, please provide the FULL docstring block including quotes/comments as per language standard.

Code:
%s

Docstring:`

	PromptReadme = `Analyze the following source code file and generate a high-level README overview section.
Focus on:
1. The purpose of this file/module.
2. Key functions/classes and their responsibilities.
3. How to use it.

Code:
%s

README Overview:`
)

// DocGenerator is responsible for generating documentation.
type DocGenerator struct {
	LLM llm.Provider
}

// NewDocGenerator creates a new documentation generator.
func NewDocGenerator(provider llm.Provider) *DocGenerator {
	return &DocGenerator{
		LLM: provider,
	}
}

// GenerateDocs generates documentation for the given code using a parser.
func (g *DocGenerator) GenerateDocs(ctx context.Context, content []byte, p parser.Parser) (string, error) {
	// Parse code into nodes
	nodes, err := p.Parse(ctx, content)
	if err != nil {
		return "", fmt.Errorf("parsing failed: %w", err)
	}

	var sb strings.Builder
	sb.WriteString("# Generated Documentation\n\n")

	for _, node := range nodes {
		fmt.Printf("Processing %s: %s...\n", node.Type, node.Name)

		docString, err := g.generateNodeDoc(ctx, node.Content)
		if err != nil {
			return "", fmt.Errorf("failed to generate doc for %s: %w", node.Name, err)
		}

		sb.WriteString(fmt.Sprintf("## %s `%s`\n\n", strings.Title(string(node.Type)), node.Name))
		sb.WriteString(docString + "\n\n")
		sb.WriteString("```" + "\n")
		sb.WriteString("```\n\n")
	}

	if len(nodes) == 0 {
		sb.WriteString("No documentable nodes found (functions/classes).\n")
		return g.generateNodeDoc(ctx, string(content))
	}

	return sb.String(), nil
}

// InjectDocstrings generates docstrings and injects them into the content.
func (g *DocGenerator) InjectDocstrings(ctx context.Context, content []byte, p parser.Parser) ([]byte, error) {
	nodes, err := p.Parse(ctx, content)
	if err != nil {
		return nil, fmt.Errorf("parsing failed: %w", err)
	}

	// Sort nodes by InsertionPoint descending so we don't mess up line numbers when inserting
	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].InsertionPoint > nodes[j].InsertionPoint
	})

	lines := bytes.Split(content, []byte("\n"))

	// Need to handle insertions. operating on slice of lines is easier.
	// But simply inserting lines works best if we go backwards.

	// Convert to []string for easier manipulation? Or stay with bytes.
	// Let's use a map or just rebuild.
	// Actually, careful with byte offsets vs line numbers.
	// The InsertionPoint is 1-indexed line number.

	for _, node := range nodes {
		fmt.Printf("Generating docstring for %s: %s...\n", node.Type, node.Name)
		doc, err := g.generateNodeDoc(ctx, node.Content)
		if err != nil {
			return nil, err
		}

		// Indent the docstring
		docLines := strings.Split(doc, "\n")
		var indentedDoc bytes.Buffer
		for _, l := range docLines {
			indentedDoc.WriteString(node.Indent)
			indentedDoc.WriteString(l)
			indentedDoc.WriteString("\n")
		}

		// Insert into content
		// InsertionPoint is 1-based index (e.g. 2 means insert after line 1, i.e. at index 1 before existing line 2)
		// Wait, if InsertionPoint is the line number *to be inserted at*, then index is InsertionPoint-1.

		idx := node.InsertionPoint - 1
		if idx < 0 {
			idx = 0
		}
		if idx > len(lines) {
			idx = len(lines)
		}

		// Insert indentedDoc
		// Go slice insertion: append(lines[:i], append(new, lines[i:]...)...)
		newLines := make([][]byte, 0)
		newLines = append(newLines, lines[:idx]...)
		newLines = append(newLines, indentedDoc.Bytes()) // Treat whole block as one "line" or split?
		// Ideally split it properly so re-parsing works if we wanted.
		// But for simple reconstruction:
		// Split indentedDoc into lines again?

		docBlockLines := bytes.Split(indentedDoc.Bytes(), []byte("\n"))
		// Remove last empty one if split creates it
		if len(docBlockLines) > 0 && len(docBlockLines[len(docBlockLines)-1]) == 0 {
			docBlockLines = docBlockLines[:len(docBlockLines)-1]
		}

		newLines = append(newLines, docBlockLines...)
		newLines = append(newLines, lines[idx:]...)

		lines = newLines
	}

	return bytes.Join(lines, []byte("\n")), nil
}

// GenerateReadme generates a README overview for the code.
func (g *DocGenerator) GenerateReadme(ctx context.Context, content []byte) (string, error) {
	prompt := fmt.Sprintf(PromptReadme, string(content))
	return g.LLM.GenerateCompletion(ctx, prompt)
}

func (g *DocGenerator) generateNodeDoc(ctx context.Context, code string) (string, error) {
	prompt := fmt.Sprintf(PromptDocString, code)
	return g.LLM.GenerateCompletion(ctx, prompt)
}
