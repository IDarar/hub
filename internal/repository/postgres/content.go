package postgres

import (
	"errors"

	"github.com/IDarar/hub/internal/domain"
	"github.com/IDarar/hub/pkg/logger"
	"gorm.io/gorm"
)

type ContentRepo struct {
	db *gorm.DB
}

func NewContentRepo(db *gorm.DB) *ContentRepo {
	return &ContentRepo{
		db: db,
	}
}
func (r *ContentRepo) Create(treatise domain.Treatise) error {
	return r.db.Create(&treatise).Error

}
func (r *ContentRepo) Delete(treatise domain.Treatise) error {
	logger.Info(treatise)
	check := r.db.Delete(&treatise).RowsAffected
	if check == 0 {
		logger.Info("Could not delete")

		return errors.New("Could not delete")
	}
	return nil
}
