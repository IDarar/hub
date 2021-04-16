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

/* TODO func (u *UsersRepo) CreateMark(domain.UserProposition, [3]interface{}) error {

	return nil
}*/
