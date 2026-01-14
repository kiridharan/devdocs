package parser

import (
	"context"
	"fmt"

	sitter "github.com/smacker/go-tree-sitter"
	"github.com/smacker/go-tree-sitter/javascript"
)

// JSParser parses JavaScript code.
type JSParser struct{}

// NewJSParser creates a new JavaScript parser.
func NewJSParser() *JSParser {
	return &JSParser{}
}

// Ensure JSParser implements Parser
var _ Parser = (*JSParser)(nil)

// Parse parses JavaScript code and returns documentable nodes.
func (p *JSParser) Parse(ctx context.Context, content []byte) ([]Node, error) {
	lang := javascript.GetLanguage()
	parser := sitter.NewParser()
	parser.SetLanguage(lang)

	tree, err := parser.ParseCtx(ctx, nil, content)
	if err != nil {
		return nil, fmt.Errorf("failed to parse javascript code: %w", err)
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
		case "function_declaration":
			nType = NodeTypeFunction
			name = getJSName(child, content)
		case "class_declaration":
			nType = NodeTypeClass
			name = getJSName(child, content)
		default:
			// TODO: Handle exported functions/classes and other patterns
			continue
		}

		// Calculate indentation
		startByte := child.StartByte()
		lineStartByte := findLineStart(content, startByte)
		indent := string(content[lineStartByte:startByte])

		// For JS, docstrings (JSDoc) typically go *before* the function/class.
		insertPoint := int(child.StartPoint().Row) + 1 // Start line (1-indexed) is where we insert (pushing down)

		nodes = append(nodes, Node{
			Type:           nType,
			Name:           name,
			Content:        child.Content(content),
			Line:           int(child.StartPoint().Row) + 1,
			InsertionPoint: insertPoint,
			Indent:         indent,
		})
	}

	return nodes, nil
}

func getJSName(node *sitter.Node, content []byte) string {
	nameNode := node.ChildByFieldName("name")
	if nameNode != nil {
		return nameNode.Content(content)
	}
	return "unknown"
}
