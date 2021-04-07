package postgres

import (
	"gorm.io/gorm"
)

type AdminsRepo struct {
	db *gorm.DB
}

func NewAdminsRepo(db *gorm.DB) *AdminsRepo {
	return &AdminsRepo{
		db: db,
	}
}

//TODO
func (r *AdminsRepo) GrantRole(id int) {
}
func (r *AdminsRepo) RevokeRole(id int) {

}
