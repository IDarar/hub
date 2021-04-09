package service

import (
	"errors"

	"github.com/IDarar/hub/internal/domain"
	"github.com/IDarar/hub/internal/repository"
)

//39.58
type AdminsService struct {
	repo repository.Admins
}

func NewAdminsService(repo repository.Admins) *AdminsService {
	return &AdminsService{
		repo: repo,
	}

}

func (u *AdminsService) GrantRole(name, role string, roles interface{}) error {

	switch true {
	case FindRole(roles.([]string), "admin"):
	case FindRole(roles.([]string), "SuperModerator"):
	default:
		return errors.New("Don't have enough rights")
	}

	return u.repo.GrantRole(name, role)
}

func FindRole(roles []string, val string) bool {
	for _, item := range roles {
		if item == val {
			return true
		}
	}
	return false
}

func (u *AdminsService) RevokeRole(user *domain.User, role string) error {

	return nil
}
