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
func (r *PropositionsRepo) Create(part domain.Proposition) error {
	logger.Error(part.TargetID)

	err := r.db.Create(&part).Error
	if err != nil {
		logger.Error(err)
		return err

	}

	return nil
}
func (r *PropositionsRepo) Delete(part domain.Proposition) error {
	logger.Info(part)
	check := r.db.Delete(&part).RowsAffected
	if check == 0 {
		logger.Info("could not delete")

		return errors.New("could not delete")
	}
	return nil
}
