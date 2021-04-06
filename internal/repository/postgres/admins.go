package postgres

import (
	"gorm.io/gorm"
)

//44.41
type AdminsRepo struct {
	db *gorm.DB
}

func NewAdminsRepo(db *gorm.DB) *UsersRepo {
	return &UsersRepo{
		db: db,
	}
}

//TODO
func (a *AdminsRepo) grantRole() {

}
func (a *AdminsRepo) revokeRole() {

}
