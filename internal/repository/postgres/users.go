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
	err := r.db.Where(&User{Name: name, Password: password}).First(&user).Error
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
	/*Db.Model(&User{}).Where("id = ?", studentId).Update("avatar", "https://forumwebappdeploytest.herokuapp.com/profileimages/"+s+".jpg")
	return err*/
	return nil
}

/* TODO func (u *UsersRepo) CreateMark(domain.UserProposition, [3]interface{}) error {

	return nil
}*/
