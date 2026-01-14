package parser

import (
	"context"
	"fmt"

	sitter "github.com/smacker/go-tree-sitter"
	"github.com/smacker/go-tree-sitter/python"
)

// PythonParser parses Python code.
type PythonParser struct{}

// NewPythonParser creates a new Python parser.
func NewPythonParser() *PythonParser {
	return &PythonParser{}
}

// Ensure PythonParser implements Parser
var _ Parser = (*PythonParser)(nil)

// Parse parses Python code and returns documentable nodes.
func (p *PythonParser) Parse(ctx context.Context, content []byte) ([]Node, error) {
	lang := python.GetLanguage()
	parser := sitter.NewParser()
	parser.SetLanguage(lang)

	tree, err := parser.ParseCtx(ctx, nil, content)
	if err != nil {
		return nil, fmt.Errorf("failed to parse python code: %w", err)
	}
	defer tree.Close()

	root := tree.RootNode()
	var nodes []Node

	// Iterate over top-level children
	count := int(root.NamedChildCount())
	for i := 0; i < count; i++ {
		child := root.NamedChild(i)
		nodeType := child.Type()

		var nType NodeType
		var name string

		switch nodeType {
		case "function_definition":
			nType = NodeTypeFunction
			name = getPythonFunctionName(child, content)
		case "class_definition":
			nType = NodeTypeClass
			name = getPythonClassName(child, content)
		default:
			continue
		}

		// Calculate indentation
		startByte := child.StartByte()
		// startRow := child.StartPoint().Row // unused
		lineStartByte := findLineStart(content, startByte)
		indent := string(content[lineStartByte:startByte])

		// Calculate insertion point (simplified: right after definition line)
		// Real implementation might need to skip existing docstrings or find exact body start.
		// For Python, it's usually inside the block, indented.
		bodyNode := child.ChildByFieldName("body")
		var insertPoint int
		var innerIndent string

		if bodyNode != nil {
			insertPoint = int(bodyNode.StartPoint().Row) + 1
			// Inner indent is usually body indent
			bodyStartByte := bodyNode.StartByte()
			bodyLineStartByte := findLineStart(content, bodyStartByte)
			innerIndent = string(content[bodyLineStartByte:bodyStartByte])
		} else {
			insertPoint = int(child.StartPoint().Row) + 2 // Fallback
			innerIndent = indent + "    "
		}

		nodes = append(nodes, Node{
			Type:           nType,
			Name:           name,
			Content:        child.Content(content),
			Line:           int(child.StartPoint().Row) + 1,
			InsertionPoint: insertPoint,
			Indent:         innerIndent,
		})
	}

	return nodes, nil
}

func getPythonFunctionName(node *sitter.Node, content []byte) string {
	// child node 'name' contains the identifier
	nameNode := node.ChildByFieldName("name")
	if nameNode != nil {
		return nameNode.Content(content)
	}
	return "unknown_function"
}

func getPythonClassName(node *sitter.Node, content []byte) string {
	nameNode := node.ChildByFieldName("name")
	if nameNode != nil {
		return nameNode.Content(content)
	}
	return "unknown_class"
}
