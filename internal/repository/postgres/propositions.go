package postgres

import (
	"errors"

	"github.com/IDarar/hub/internal/domain"
	"github.com/IDarar/hub/pkg/logger"
	"gorm.io/gorm"
)

type PropositionsRepo struct {
	db *gorm.DB
}

func NewPropositionsRepo(db *gorm.DB) *PropositionsRepo {
	return &PropositionsRepo{
		db: db,
	}
}

/*err := r.db.Model(&treatise).Association("Parts").Append([]domain.Part{part})
if err != nil {
	logger.Error(err)
	return err

}*/
func (r *PropositionsRepo) Create(proposition domain.Proposition) error {
	err := r.db.Model(proposition).First(&proposition).Error
	if err == nil {
		logger.Error("found")
		return errors.New("proposition already exists")
	}
	treatise := &domain.Treatise{ID: proposition.TargetID}

	err = r.db.Model(&treatise).Association("Propositions").Append(&proposition)
	if err == nil {
		logger.Error(err)
		return err
	}
	err = r.db.Model(proposition).First(&proposition).Error
	if err == nil {
		logger.Info("found")
		return nil
	}
	part := &domain.Part{ID: proposition.TargetID}
	err = r.db.Model(&part).Association("Propositions").Append(&proposition)
	if err != nil {
		logger.Error(err)
		return err
	}
	err = r.db.Model(proposition).First(&proposition).Error
	if err != nil {
		logger.Error("found")
		return err
	}
	return nil
}

func (r *PropositionsRepo) Delete(proposition domain.Proposition) error {
	logger.Info(proposition)
	check := r.db.Delete(&proposition).RowsAffected
	if check == 0 {
		logger.Info("could not delete")

		return errors.New("could not delete")
	}

	return nil
}

func (r *PropositionsRepo) GetByID(id string) (domain.Proposition, error) {
	var pr domain.Proposition
	err := r.db.Where(&domain.Proposition{ID: id}).First(&pr).Error
	if err != nil {
		logger.Warn("not found proposition ", err)
		return domain.Proposition{}, err

	}

	return pr, nil
}
func (r *PropositionsRepo) Update(prop domain.Proposition, createReferences, deleteReferences []string) error {
	logger.Info(prop)

	err := r.db.Model(&prop).Updates(&prop).Error
	if err != nil {
		logger.Error(err)
		return err
	}
	ref := domain.Reference{}
	if len(createReferences) != 0 {

		for _, v := range createReferences {
			r.db.Where(&domain.Reference{Target: prop.ID}).First(&ref)
			ref.TargetProposition = v
			r.db.Create(&ref)
		}
	}

	return nil
}
