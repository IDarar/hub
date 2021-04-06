package service

import (
	"hub/internal/domain"
	"hub/internal/repository"
)

//39.58
type Users struct {
	repo repository.Users
}

func NewUsersService(repo repository.Users) *Users {
	return &Users{
		repo: repo,
	}
}

func (u *Users) CreateSession(domain.UserProposition, [3]interface{}) error {

	return nil

}

func (u *Users) CreateMark(domain.UserProposition, [3]interface{}) error {

	return nil
}
