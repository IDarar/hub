package postgres

import (
	"context"

	"github.com/IDarar/hub/internal/domain"
	"github.com/IDarar/hub/pkg/logger"

	"gorm.io/gorm"
)

//44.41
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

/* TODO func (u *UsersRepo) CreateMark(domain.UserProposition, [3]interface{}) error {

	return nil
}*/
