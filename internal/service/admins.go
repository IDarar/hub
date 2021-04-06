package service

import (
	"hub/internal/domain"
	"hub/internal/repository"
)

//39.58
type Admins struct {
	repo repository.Admins
}

func NewAdminsService(repo repository.Users) *Users {
	return &Users{
		repo: repo,
	}
}

//TODO think about passing User's modls or IDs, know what context is used for
//TODO know should service methods and repo ones have identic description (what vars they get)
func (u *Users) grantRole(user *domain.User, role string) error {

	return nil
}
func (u *Users) revokeRole(user *domain.User, role string) error {

	return nil
}
