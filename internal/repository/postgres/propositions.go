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
