package embedder

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/firebase/genkit/go/ai"
)

type OllamaEmbedder struct {
	Model string
}

func (e *OllamaEmbedder) Name() string {
	return "ollama:" + e.Model
}

func (e *OllamaEmbedder) Embed(ctx context.Context, req *ai.EmbedRequest) (*ai.EmbedResponse, error) {
	// 回傳的格式
	respEmbeds := &ai.EmbedResponse{
		Embeddings: make([]*ai.DocumentEmbedding, len(req.Documents)),
	}

	for i, doc := range req.Documents {
		if len(doc.Content) == 0 {
			return nil, fmt.Errorf("document %d has no content", i)
		}
		text := doc.Content[0].Text

		payload := map[string]any{
			"model":  e.Model,
			"prompt": text,
		}
		body, err := json.Marshal(payload)
		if err != nil {
			return nil, err
		}

		httpReq, err := http.NewRequestWithContext(ctx, "POST", "http://localhost:11434/api/embeddings", bytes.NewReader(body))
		if err != nil {
			return nil, err
		}
		httpReq.Header.Set("Content-Type", "application/json")

		resp, err := http.DefaultClient.Do(httpReq)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		var result struct {
			Embedding []float32 `json:"embedding"`
		}
		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			return nil, err
		}

		respEmbeds.Embeddings[i] = &ai.DocumentEmbedding{Embedding: result.Embedding}
	}

	return respEmbeds, nil
}
