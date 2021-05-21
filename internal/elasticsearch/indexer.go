package elasticsearch

import (
	"context"

	"github.com/IDarar/hub/internal/domain"
	esv7 "github.com/elastic/go-elasticsearch/v7"
)

type Content interface {
	Index(ctx context.Context, content domain.Treatise) error
}
type Requst interface {
	Index(ctx context.Context, req IndexedRequest) error
}
type Indexers struct {
	Content Content
	Request Requst
}

func NewIndexer(client *esv7.Client) *Indexers {
	return &Indexers{
		Content: NewContent(client),
		Request: NewRequest(client),
	}
}
