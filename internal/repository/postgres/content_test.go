package postgres

import (
	"testing"

	"github.com/IDarar/hub/internal/domain"
	"github.com/IDarar/hub/pkg/logger"
	"github.com/stretchr/testify/assert"
)

func TestContent_Create(t *testing.T) {
	db, truncate := NewTestPostgres(t)
	if db == nil {
		logger.Info("IS NIL")
	}
	contentRepo := NewContentRepo(db)

	defer truncate("part_propositions", "treatise_propositions", "propositions", "treatises", "parts")
	id := "string"

	err := contentRepo.Create(domain.Treatise{ID: id})
	assert.Error(t, err)

	tr, err := contentRepo.GetByID(id)
	assert.NoError(t, err)
	assert.NotNil(t, tr)

}
