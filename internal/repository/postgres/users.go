package postgres

import (
	"hub/internal/domain"

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
func GetUserByID(int) (*domain.User, error) {
	return &domain.User{}, nil
}
func (u *UsersRepo) CreateMark(domain.UserProposition, [3]interface{}) error {

	return nil
}
