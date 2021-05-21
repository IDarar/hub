package elasticsearch

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"time"

	"github.com/IDarar/hub/pkg/logger"
	"github.com/google/uuid"

	esv7 "github.com/elastic/go-elasticsearch/v7"
	esv7api "github.com/elastic/go-elasticsearch/v7/esapi"
)

type RequestIndexer struct {
	client *esv7.Client
	index  string
}

type IndexedRequest struct {
	ID     string    `json:"id,omitempty"`
	Url    string    `json:"url,omitempty"`
	Method string    `json:"method,omitempty"`
	Date   time.Time `json:"date,omitempty"`
}

func NewRequest(client *esv7.Client) *RequestIndexer {
	return &RequestIndexer{
		client: client,
		index:  "request",
	}
}
func (i *RequestIndexer) Index(ctx context.Context, req IndexedRequest) error {

	body := IndexedRequest{
		ID:     uuid.NewString(),
		Url:    req.Url,
		Method: req.Method,
		Date:   time.Now(),
	}

	var buf bytes.Buffer

	if err := json.NewEncoder(&buf).Encode(body); err != nil {
		logger.Error(err)
		return err
	}

	eReq := esv7api.IndexRequest{
		Index:      i.index,
		Body:       &buf,
		DocumentID: req.ID,
		Refresh:    "true",
	}

	resp, err := eReq.Do(ctx, i.client)
	if err != nil {
		logger.Error(err)
		return err
	}
	defer resp.Body.Close()

	if resp.IsError() {
		logger.Error(err)
		return err
	}

	io.Copy(ioutil.Discard, resp.Body)

	return nil
}
