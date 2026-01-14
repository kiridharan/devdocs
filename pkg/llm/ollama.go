package llm

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

// OllamaProvider implements the Provider interface for Ollama.
type OllamaProvider struct {
	BaseURL string
	Model   string
	Client  *http.Client
}

// Ensure OllamaProvider implements Provider
var _ Provider = (*OllamaProvider)(nil)

// NewOllamaProvider creates a new Ollama provider.
func NewOllamaProvider(baseURL, model string) *OllamaProvider {
	if baseURL == "" {
		baseURL = "http://localhost:11434"
	}
	if model == "" {
		model = "llama3" // Default fallback
	}
	return &OllamaProvider{
		BaseURL: baseURL,
		Model:   model,
		Client:  &http.Client{Timeout: 60 * time.Second},
	}
}

// GenerateRequest represents the request body for Ollama API.
type GenerateRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
	Stream bool   `json:"stream"`
}

// GenerateResponse represents the response body from Ollama API.
type GenerateResponse struct {
	Response string `json:"response"`
	Done     bool   `json:"done"`
}

// GenerateCompletion generates a completion using Ollama.
func (p *OllamaProvider) GenerateCompletion(ctx context.Context, prompt string) (string, error) {
	reqBody := GenerateRequest{
		Model:  p.Model,
		Prompt: prompt,
		Stream: false,
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", p.BaseURL+"/api/generate", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := p.Client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send request to Ollama: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("ollama returned non-200 status: %s", resp.Status)
	}

	var genResp GenerateResponse
	if err := json.NewDecoder(resp.Body).Decode(&genResp); err != nil {
		return "", fmt.Errorf("failed to decode response: %w", err)
	}

	return strings.TrimSpace(genResp.Response), nil
}
