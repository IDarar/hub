package service

import (
	"context"
	"time"

	"github.com/IDarar/hub/internal/domain"
	"github.com/IDarar/hub/internal/repository"
)

//39.58
type UserService struct {
	repo repository.Users
}

func NewUsersService(repo repository.Users) *UserService {
	return &UserService{
		repo: repo,
	}
}
func (s *UserService) SignUp(ctx context.Context, input SignUpInput) error {
	user := domain.User{
		Name: input.Name,
		//Password:     s.hasher.Hash(input.Password),
		Email:        input.Email,
		RegisteredAt: time.Now(),
		LastVisitAt:  time.Now(),
	}

	if err := s.repo.Create(ctx, user); err != nil {
		return err
	}
	return nil

}

func (u *UserService) CreateMark(domain.UserProposition, [3]interface{}) error {

	return nil
}
