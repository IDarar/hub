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
func (r *PropositionsRepo) Update(prop domain.Proposition,
	createReferences, deleteReferences []string,
	createNotes, deleteNotes []domain.Note) error {
	logger.Info(prop)

	err := r.db.Model(&prop).Updates(&prop).Error
	if err != nil {
		logger.Error(err)
		return err
	}
	if len(createReferences) != 0 {

		for _, v := range createReferences {
			ref := domain.Reference{}
			err = r.db.Where(&domain.Reference{Target: prop.ID, TargetProposition: v}).First(&ref).Error
			if err == nil {
				logger.Error("ref already exists")
				return errors.New("ref already exists")
			}
			ref.Target = prop.ID
			ref.TargetProposition = v

			err = r.db.Create(&ref).Error
			if err != nil {
				logger.Error(err)
				return errors.New("could not create")
			}
		}
	}
	if len(deleteReferences) != 0 {

		for _, v := range deleteReferences {
			ref := domain.Reference{}
			err = r.db.Where(&domain.Reference{Target: prop.ID, TargetProposition: v}).First(&ref).Error
			if err != nil {
				logger.Error("ref don't exist")
				return errors.New("ref don't exist")
			}
			ref.Target = prop.ID
			ref.TargetProposition = v

			count := r.db.Delete(&ref).RowsAffected
			if count == 0 {
				logger.Error("could not delete")
				return errors.New("could not delete")
			}
		}
	}
	if len(createNotes) != 0 {

		for _, v := range createNotes {
			note := domain.Note{}
			err = r.db.Where(&domain.Note{Target: prop.ID, Text: v.Text}).First(&note).Error
			if err == nil {
				logger.Error("note already exists")
				return errors.New("note already exists")
			}
			note.Target = prop.ID
			note.TreatiseID = prop.TargetID
			note.Text = v.Text
			note.Type = v.Type
			err = r.db.Create(&note).Error
			if err != nil {
				logger.Error(err)
				return errors.New("could not create")
			}
		}
	}
	if len(deleteNotes) != 0 {

		for _, v := range deleteNotes {
			note := domain.Note{}
			err = r.db.Where(&domain.Note{Target: prop.ID, Text: v.Text}).First(&note).Error
			if err != nil {
				logger.Error("note don't exist")
				return errors.New("note don't exist")
			}
			note.Target = prop.ID

			count := r.db.Delete(&note).RowsAffected
			if count == 0 {
				logger.Error("could not delete")
				return errors.New("could not delete")
			}
		}
	}
	return nil
}
