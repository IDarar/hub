package postgres

import (
	"context"
	"errors"

	"github.com/IDarar/hub/internal/domain"
	"github.com/IDarar/hub/pkg/logger"

	"gorm.io/gorm"
)

type UsersRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UsersRepo {
	return &UsersRepo{
		db: db,
	}
}
func (r *UsersRepo) Create(ctx context.Context, user domain.User) error {
	err := r.db.Create(&user).Error

	return err
}
func (r *UsersRepo) GetByCredentials(ctx context.Context, name, password string) (domain.User, error) {
	var user domain.User
	err := r.db.Where(&domain.User{Name: name, Password: password}).First(&user).Error
	if err != nil {
		logger.Warn("Not found user ", err)
		return domain.User{}, err

	}

	return user, nil
}
func (r *UsersRepo) GetByName(name string) (domain.User, error) {
	var user domain.User
	err := r.db.Where(&domain.User{Name: name}).First(&user).Error
	if err != nil {
		logger.Warn("Not found user ", err)
		return domain.User{}, err

	}

	return user, nil
}
func (r *UsersRepo) GetUserByID(int) (*domain.User, error) {
	return &domain.User{}, nil
}
func (r *UsersRepo) SetSession(userId int, session domain.Session) error {
	//"refresh_token", session.RefreshToken
	err := r.db.Model(&domain.Session{}).Where("user_id = ?", userId).Updates(domain.Session{RefreshToken: session.RefreshToken, ExpiresAt: session.ExpiresAt}).Error
	return err
}
func (r *UsersRepo) GetRoleByID(userId int) ([]string, error) {
	user := domain.User{ID: userId}
	var roles []string
	rolesDB := []domain.Role{}

	err := r.db.Model(&user).Association("Roles").Find(&rolesDB)
	if err != nil {
		logger.Warn("User does not have roles", err)
		return roles, err
	}
	for _, v := range rolesDB {
		roles = append(roles, v.Role)
	}

	logger.Info(roles, " from slice")
	return roles, err
}

//There are no foreign keys for target
//It also don't check the target of object
//Target should exist from input
func (r *UsersRepo) AddTreatise(tr domain.UserTreatise) error {

	tr.Status = "In progress"
	return r.db.Create(&tr).Error
}
func (r *UsersRepo) UpdateTreatise(tr domain.UserTreatise) error {
	logger.Info(tr)
	if tr.IsCompleted == nil {
		logger.Info("is nil")
	}

	err := r.db.Model(&tr).Updates(&tr).Error
	if err != nil {
		logger.Error(err)
		return err
	}

	return nil
}
func (r *UsersRepo) AddProposition(pr domain.UserProposition) error {

	err := r.db.Model(pr).First(&pr).Error
	logger.Info(pr)
	if err == nil {
		logger.Error("found")
		return errors.New("already added")
	}
	return r.db.Create(&pr).Error
}
func (r *UsersRepo) UpdateProposition(pr domain.UserProposition) error {
	logger.Info(pr)
	if pr.IsCompleted == nil {
		logger.Info("is nil")
	}

	err := r.db.Model(&pr).Updates(&pr).Error
	if err != nil {
		logger.Error(err)
		return err
	}

	return nil
}
func (r *UsersRepo) AddPart(part domain.UserPart) error {
	err := r.db.Model(part).First(&part).Error
	logger.Info(part)
	if err == nil {
		logger.Error("found")
		return errors.New("already added")
	}
	return r.db.Create(&part).Error
}
func (r *UsersRepo) UpdatePart(part domain.UserPart) error {
	logger.Info(part)
	if part.IsCompleted == nil {
		logger.Info("is nil")
	}

	err := r.db.Model(&part).Updates(&part).Error
	if err != nil {
		logger.Error(err)
		return err
	}

	return nil
}

//IT does not check if usertr exists so the proper value should be sent
func (r *UsersRepo) RateTreatise(tr domain.UserTreatise, rate domain.Rate) error {
	logger.Info("RATE ", rate)

	err := r.db.Where("user_id = ? AND type = ? AND target_id = ?",
		rate.UserID,
		rate.Type,
		rate.TargetID).First(&rate).Error

	if err == nil {
		logger.Info("exists ", rate)
		logger.Error(err)
		err = r.db.Model(&tr).Updates(&tr).Error
		if err != nil {
			logger.Error(err)
			return err
		}
		err = r.db.Model(&rate).Updates(&rate).Error
		if err != nil {
			logger.Error(err)
			return err
		}
		return nil
	}
	logger.Info("RATE ", rate)

	err = r.db.Model(&domain.Treatise{ID: rate.TargetID}).Association("Rates").Append(&rate)
	if err != nil {
		logger.Error(err)
		return err
	}

	err = r.db.Model(&domain.UserLists{UserID: tr.UserID}).Association("Rates").Append(&rate)
	if err != nil {
		logger.Error(err)
		return err
	}

	err = r.db.Model(&tr).Updates(&tr).Error
	if err != nil {
		logger.Error(err)
		return err
	}
	return nil
}

//IT does not check if userpart exists so the proper value should be sent
func (r *UsersRepo) RatePart(part domain.UserPart, rate domain.Rate) error {
	logger.Info("RATE ", rate)

	err := r.db.Where("user_id = ? AND type = ? AND target_id = ?",
		rate.UserID,
		rate.Type,
		rate.TargetID).First(&rate).Error

	if err == nil {
		logger.Info("exists ", rate)
		logger.Error(err)
		err = r.db.Model(&part).Updates(&part).Error
		if err != nil {
			logger.Error(err)
			return err
		}
		err = r.db.Model(&rate).Updates(&rate).Error
		if err != nil {
			logger.Error(err)
			return err
		}
		return nil
	}
	logger.Info("RATE ", rate)

	err = r.db.Model(&domain.Part{ID: rate.TargetID}).Association("Rates").Append(&rate)
	if err != nil {
		logger.Error(err)
		return err
	}

	err = r.db.Model(&domain.UserLists{UserID: part.UserID}).Association("Rates").Append(&rate)
	if err != nil {
		logger.Error(err)
		return err
	}

	err = r.db.Model(&part).Updates(&part).Error
	if err != nil {
		logger.Error(err)
		return err
	}
	return nil
}
func (r *UsersRepo) RateProposition(pr domain.UserProposition, rate domain.Rate) error {
	logger.Info("RATE ", rate)

	err := r.db.Where("user_id = ? AND type = ? AND target_id = ?",
		rate.UserID,
		rate.Type,
		rate.TargetID).First(&rate).Error

	if err == nil {
		logger.Info("exists ", rate)
		logger.Error(err)
		err = r.db.Model(&pr).Updates(&pr).Error
		if err != nil {
			logger.Error(err)
			return err
		}
		err = r.db.Model(&rate).Updates(&rate).Error
		if err != nil {
			logger.Error(err)
			return err
		}
		return nil
	}
	logger.Info("RATE ", rate)

	err = r.db.Model(&domain.Proposition{ID: rate.TargetID}).Association("Rates").Append(&rate)
	if err != nil {
		logger.Error(err)
		return err
	}

	err = r.db.Model(&domain.UserLists{UserID: pr.UserID}).Association("Rates").Append(&rate)
	if err != nil {
		logger.Error(err)
		return err
	}

	err = r.db.Model(&pr).Updates(&pr).Error
	if err != nil {
		logger.Error(err)
		return err
	}
	return nil

}
func (r *UsersRepo) DeleteRateTreatise(tr domain.UserTreatise, rate domain.Rate) error {
	logger.Info("RATE ", rate)
	logger.Warn("TRU ", tr)
	err := r.db.Where("user_id = ? AND type = ? AND target_id = ?",
		rate.UserID,
		rate.Type,
		rate.TargetID).First(&rate).Error
	if err != nil {
		logger.Error(err)
		return err
	}
	check := r.db.Delete(rate).RowsAffected
	if check == 0 {
		logger.Error(errors.New("not deleted, probably object does not exist"))
		return errors.New("not deleted, probably object does not exist")
	}
	logger.Info("TR ", tr.ImportanceRate)
	rateToDel, ent := chekType(tr)
	logger.Info(ent)
	err = r.db.Model(&tr).Select(rateToDel).Updates(ent).Error
	if err != nil {
		logger.Error(err)
		return err
	}
	return nil

}
func (r *UsersRepo) DeleteRatePart(part domain.UserPart, rate domain.Rate) error {
	logger.Info("RATE ", rate)
	logger.Warn("TRU ", part)
	err := r.db.Where("user_id = ? AND type = ? AND target_id = ?",
		rate.UserID,
		rate.Type,
		rate.TargetID).First(&rate).Error
	if err != nil {
		logger.Error(err)
		return err
	}
	check := r.db.Delete(rate).RowsAffected
	if check == 0 {
		logger.Error(errors.New("not deleted, probably object does not exist"))
		return errors.New("not deleted, probably object does not exist")
	}
	logger.Info("TR ", part.ImportanceRate)
	rateToDel, ent := chekType(part)
	logger.Info(ent)
	err = r.db.Model(&part).Select(rateToDel).Updates(ent).Error
	if err != nil {
		logger.Error(err)
		return err
	}
	return nil
}
func (r *UsersRepo) DeleteRateProposition(pr domain.UserProposition, rate domain.Rate) error {
	logger.Info("RATE ", rate)
	logger.Warn("TRU ", pr)
	err := r.db.Where("user_id = ? AND type = ? AND target_id = ?",
		rate.UserID,
		rate.Type,
		rate.TargetID).First(&rate).Error
	if err != nil {
		logger.Error(err)
		return err
	}
	check := r.db.Delete(rate).RowsAffected
	if check == 0 {
		logger.Error(errors.New("not deleted, probably object does not exist"))
		return errors.New("not deleted, probably object does not exist")
	}
	logger.Info("TR ", pr.ImportanceRate)
	rateToDel, ent := chekType(pr)
	logger.Info(ent)
	err = r.db.Model(&pr).Select(rateToDel).Updates(ent).Error
	if err != nil {
		logger.Error(err)
		return err
	}
	return nil
}
func chekType(ent interface{}) (string, interface{}) {
	if part, ok := ent.(domain.UserPart); ok {
		if part.DifficultyRate != 0 {
			part.DifficultyRate = 0
			return "difficulty_rate", part
		}
		if part.ImportanceRate != 0 {
			part.ImportanceRate = 0
			return "importance_rate", part
		}
		if part.InconsistencyRate != 0 {
			part.InconsistencyRate = 0
			return "inconsistency_rate", part
		}
	}
	if tr, ok := ent.(domain.UserTreatise); ok {
		if tr.DifficultyRate != 0 {
			tr.DifficultyRate = 0
			return "difficulty_rate", tr
		}
		if tr.ImportanceRate != 0 {
			tr.ImportanceRate = 0
			return "importance_rate", tr
		}
		if tr.InconsistencyRate != 0 {
			tr.InconsistencyRate = 0
			return "inconsistency_rate", tr
		}
	}
	if pr, ok := ent.(domain.UserProposition); ok {
		if pr.DifficultyRate != 0 {
			pr.DifficultyRate = 0
			return "difficulty_rate", pr
		}
		if pr.ImportanceRate != 0 {
			pr.ImportanceRate = 0
			return "importance_rate", pr
		}
		if pr.InconsistencyRate != 0 {
			pr.InconsistencyRate = 0
			return "inconsistency_rate", pr
		}
	}
	return "invalid type", ent
}
