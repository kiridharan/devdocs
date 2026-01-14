package llm

import "context"

// Global constants for providers
const (
	ProviderOllama = "ollama"
	ProviderOpenAI = "openai"
)

// Provider defines the interface for an LLM provider.
type Provider interface {
	// GenerateCompletion sends a prompt to the LLM and returns the completion.
	GenerateCompletion(ctx context.Context, prompt string) (string, error)
}
