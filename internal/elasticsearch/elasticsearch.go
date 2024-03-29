package elasticsearch

import (
	"fmt"

	"github.com/IDarar/hub/internal/config"
	esv7 "github.com/elastic/go-elasticsearch/v7"
)

// NewElasticSearch instantiates the ElasticSearch client using configuration defined in environment variables.
func NewElasticSearch(conf config.Config) (es *esv7.Client, err error) {
	es, err = esv7.NewDefaultClient()
	if err != nil {
		return nil, fmt.Errorf("elasticsearch.Open %w", err)
	}

	res, err := es.Info()
	if err != nil {
		return nil, fmt.Errorf("es.Info %w", err)
	}
	defer func() {
		err = res.Body.Close()
	}()

	return es, nil
}
