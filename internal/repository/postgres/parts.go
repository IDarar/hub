package postgres

import (
	"errors"

	"github.com/IDarar/hub/internal/domain"
	"github.com/IDarar/hub/pkg/logger"
	"gorm.io/gorm"
)

type PartsRepo struct {
	db *gorm.DB
}

func NewPartsRepo(db *gorm.DB) *PartsRepo {
	return &PartsRepo{
		db: db,
	}
}

/*err := r.db.Model(&treatise).Association("Parts").Append([]domain.Part{part})
if err != nil {
	logger.Error(err)
	return err

}*/
func (r *PartsRepo) Create(part domain.Part) error {
	logger.Error(part.TargetID)

	err := r.db.Create(&part).Error
	if err != nil {
		logger.Error(err)
		return err

	}

	return nil
}
func (r *PartsRepo) Update(part domain.Part, createLiterature, deleteLiterature []string) error {
	logger.Info(part)

	err := r.db.Model(&part).Updates(&part).Error
	if err != nil {
		logger.Error(err)
		return err
	}

	if len(createLiterature) != 0 {

		for _, v := range createLiterature {
			lit := domain.Literature{}
			err = r.db.Where(&domain.Literature{TargetID: part.ID, Title: v}).First(&lit).Error
			if err == nil {
				logger.Error("literature already exists")
				return errors.New("literature already exists")
			}
			lit.TargetID = part.ID
			lit.Title = v

			err = r.db.Create(&lit).Error
			if err != nil {
				logger.Error(err)
				return errors.New("could not create")
			}
		}
	}
	if len(deleteLiterature) != 0 {

		for _, v := range deleteLiterature {
			lit := domain.Literature{}
			err = r.db.Where(&domain.Literature{TargetID: part.ID, Title: v}).First(&lit).Error
			if err != nil {
				logger.Error("literature don't exist")
				return errors.New("literature don't exist")
			}
			lit.TargetID = part.ID
			lit.Title = v

			count := r.db.Delete(&lit).RowsAffected
			if count == 0 {
				logger.Error("could not delete")
				return errors.New("could not delete")
			}
		}
	}
	return nil
}

func (r *PartsRepo) Delete(part domain.Part) error {
	logger.Info(part)
	props := []*domain.Proposition{}

	r.db.Model(&part).Association("Propositions").Find(&props)
	logger.Info("props ", props)

	r.db.Delete(&props)
	result := r.db.Delete(&part).RowsAffected
	if result == 0 {
		logger.Error("not deleted ", result)
		return errors.New("not deleted part, probably it does not exist")
	}

	return nil
}
