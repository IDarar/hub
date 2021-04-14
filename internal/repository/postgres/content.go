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
	err := r.db.Model(treatise).First(&treatise).Error
	if err == nil {
		logger.Error("found")
		return errors.New("treatise already exists")
	}
	return r.db.Create(&treatise).Error

}
func (r *ContentRepo) Update(treatise domain.Treatise) error {
	logger.Info(treatise)

	err := r.db.Model(&treatise).Updates(&treatise).Error
	if err != nil {
		logger.Error(err)
		return err
	}

	return nil
}
func (r *ContentRepo) Delete(treatise domain.Treatise) error {
	logger.Info(treatise)
	check := r.db.Delete(&treatise).RowsAffected
	if check == 0 {
		logger.Info("could not delete")

		return errors.New("could not delete")
	}
	return nil
}
func (r *ContentRepo) GetByID(id string) (domain.Treatise, error) {
	var tr domain.Treatise
	err := r.db.Where(&domain.Treatise{ID: id}).First(&tr).Error
	if err != nil {
		logger.Warn("Not found treatise ", err)
		return domain.Treatise{}, err

	}

	return tr, nil
}
