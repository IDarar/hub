package postgres

import (
	"github.com/IDarar/hub/internal/domain"
	"github.com/IDarar/hub/pkg/logger"
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
func (r *AdminsRepo) GrantRole(name, role string) error {

	user, err := r.GetByName(name)
	if err != nil {
		logger.Error(err)
		return err
	}
	roleStr := domain.Role{Role: role, Users: []domain.User{{ID: user.ID}}}

	return r.db.Model(&roleStr).Association("Users").Append([]domain.User{})

}
func (r *AdminsRepo) RevokeRole(id int) {

}

func (r *AdminsRepo) GetByName(name string) (domain.User, error) {
	var user domain.User
	err := r.db.Where(&domain.User{Name: name}).First(&user).Error
	if err != nil {
		logger.Warn("Not found user ", err)
		return domain.User{}, err

	}
	return user, nil
}
