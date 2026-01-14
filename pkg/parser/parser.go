package parser

import (
	"context"
)

// NodeType represents the type of a code node.
type NodeType string

const (
	NodeTypeFunction NodeType = "function"
	NodeTypeClass    NodeType = "class"
	NodeTypeMethod   NodeType = "method"
)

// Node represents a code structure to be documented.
type Node struct {
	Type           NodeType
	Name           string
	Content        string
	Line           int
	InsertionPoint int    // Line number where docstring should be inserted (0-indexed or 1-indexed? Let's use 1-indexed to match Line)
	Indent         string // Indentation string to prefix docstring lines
}

// Parser defines the interface for language-specific parsers.
type Parser interface {
	// Parse parses the source code and returns a list of documentable nodes.
	Parse(ctx context.Context, content []byte) ([]Node, error)
}

func findLineStart(content []byte, idx uint32) uint32 {
	for i := int(idx) - 1; i >= 0; i-- {
		if content[i] == '\n' {
			return uint32(i + 1)
		}
	}
	return 0
}
