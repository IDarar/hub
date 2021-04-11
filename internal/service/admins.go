package service

import (
	"errors"

	"github.com/IDarar/hub/internal/domain"
	"github.com/IDarar/hub/internal/repository"
	"github.com/IDarar/hub/pkg/logger"
)

//39.58
type AdminsService struct {
	repo  repository.Admins
	users User
}

func NewAdminsService(repo repository.Admins, users User) *AdminsService {
	return &AdminsService{
		repo:  repo,
		users: users,
	}

}

//Roles should be passed by rights desc
func checkRigths(roles interface{}, enoughRole ...string) error {

	for i := range enoughRole {
		if FindRole(roles.([]string), enoughRole[i]) {
			return nil
		}
	}

	return errors.New("don't have enough rights")
}
func FindRole(roles []string, val string) bool {
	for _, item := range roles {
		if item == val {

			return true
		}
	}
	return false
}

func (s *AdminsService) GrantRole(name, role string, roles interface{}) error {
	roles, err := s.users.GetRoleById(roles.(int))
	if err != nil {
		logger.Error(err)
		return err
	}

	switch true {
	case FindRole(roles.([]string), "admin"):
	case FindRole(roles.([]string), "SuperModerator"):
		if role == "admin" || role == "SuperModerator" {
			return errors.New("don't have enough rights")
		}
	default:
		return errors.New("don't have enough rights")
	}

	return s.repo.GrantRole(name, role)
}

func (s *AdminsService) RevokeRole(user *domain.User, role string) error {
	return nil
}
