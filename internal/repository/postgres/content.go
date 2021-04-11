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
func (r *ContentRepo) Update(treatise domain.Treatise) error {
	logger.Info(treatise)
	//c := r.db.Save(&treatise).RowsAffected
	err := r.db.Model(&treatise).Updates(&treatise).Error
	if err != nil {
		logger.Error(err)
		return err
	}
	/*logger.Error("COUNT ")

	err = r.db.Save(&treatise).Error
	if err != nil {
		logger.Error(err)
		return err
	}








		err := r.db.Model(&treatise).Updates(domain.Treatise{
		Date:        treatise.Date,
		Description: treatise.Description,
		Title:       treatise.Title,
	}).Error
	if err != nil {
		logger.Error(err)
		return err
	}




	*/
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
