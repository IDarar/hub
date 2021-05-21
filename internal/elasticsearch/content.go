package elasticsearch

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"time"

	"github.com/IDarar/hub/internal/domain"
	esv7 "github.com/elastic/go-elasticsearch/v7"
	esv7api "github.com/elastic/go-elasticsearch/v7/esapi"
)

type ContentIndexer struct {
	client *esv7.Client
	index  string
}

type indexedContent struct {
	ID          string    `json:"id,omitempty"`
	Title       string    `json:"title,omitempty"`
	Description string    `json:"description,omitempty"`
	Date        time.Time `json:"date,omitempty"`
}

func NewContent(client *esv7.Client) *ContentIndexer {
	return &ContentIndexer{
		client: client,
		index:  "content",
	}
}
func (i *ContentIndexer) Index(ctx context.Context, content domain.Treatise) error {

	body := indexedContent{
		ID:          content.ID,
		Description: content.Description,
		Title:       content.Title,
		Date:        time.Now(),
	}

	var buf bytes.Buffer

	if err := json.NewEncoder(&buf).Encode(body); err != nil {
		return err
	}

	req := esv7api.IndexRequest{
		Index:      i.index,
		Body:       &buf,
		DocumentID: content.ID,
		Refresh:    "true",
	}

	resp, err := req.Do(ctx, i.client)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.IsError() {
		return err
	}

	io.Copy(ioutil.Discard, resp.Body)

	return nil
}
