package postgres

import (
	"context"

	"github.com/IDarar/hub/internal/domain"

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

func (r *UsersRepo) GetUserByID(int) (*domain.User, error) {
	return &domain.User{}, nil
}

/* TODO func (u *UsersRepo) CreateMark(domain.UserProposition, [3]interface{}) error {

	return nil
}*/
