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
	//this
	logger.Info(treatise)

	err := r.db.Where("target_id = ?", treatise.ID).Delete(&domain.Proposition{TargetID: treatise.ID}).Error
	if err != nil {
		logger.Error("deleting treatise's props", err)
	}

	parts := []*domain.Part{}

	props := []*domain.Proposition{}

	propsToDel := []*domain.Proposition{}

	err = r.db.Model(&treatise).Association("Parts").Find(&parts)

	for i := 0; i < len(parts); i++ {
		err = r.db.Model(&parts[i]).Association("Propositions").Find(&props)
		for i := 0; i < len(props); i++ {
			propsToDel = append(propsToDel, props[i])
			logger.Info("PROPS  ", props[i].ID)
		}
	}
	logger.Info("propsToDel", propsToDel)
	for i := 0; i < len(props); i++ {
		logger.Info("PROPS  ", props[i].ID)
	}
	//logger.Info("PROPS  ", props)

	if err != nil {
		logger.Error(err)
		return err
	}

	err = r.db.Where("target_id = ?", treatise.ID).Delete(&domain.Proposition{TargetID: treatise.ID}).Error
	if err != nil {
		logger.Error(err)
		return err
	}

	logger.Info("PARTS ", parts)
	r.db.Delete(&propsToDel)
	r.db.Delete(&parts)
	result := r.db.Delete(&treatise).RowsAffected
	if result == 0 {
		logger.Error("not deleted ", result)
		return errors.New("not deleted proposition, probably it does not exist")
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
