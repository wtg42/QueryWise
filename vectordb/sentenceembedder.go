package vectordb

import (
	"context"

	"github.com/firebase/genkit/go/ai"
)

type SentenceTransformerEmbedder struct {
	APIEndpoint string
}

func (s *SentenceTransformerEmbedder) Name() string {
	return "paraphrase-multilingual-mpnet-base-v2"
}

func (s *SentenceTransformerEmbedder) Embed(ctx context.Context, req *ai.EmbedRequest) (*ai.EmbedResponse, error) {
	// do something like request to embeding model server...
	return nil, nil
}
