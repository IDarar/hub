package service

import (
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

//TODO think about passing User's modls or IDs, know what context is used for
//TODO know should service methods and repo ones have identic description (what vars they get)
func (u *AdminsService) grantRole(user *domain.User, role string) error {

	return nil
}
func (u *AdminsService) revokeRole(user *domain.User, role string) error {

	return nil
}
