package service

import (
	"context"

	"github.com/IDarar/hub/internal/domain"
	"github.com/IDarar/hub/internal/repository"
)

type SignUpInput struct {
	Name     string
	Email    string
	Password string
}

//all interfaces there are described
type User interface {
	SignUp(ctx context.Context, input SignUpInput) error

	CreateMark(domain.UserProposition, [3]interface{}) error
}
type Admin interface {
	grantRole(user *domain.User, role string) error
	revokeRole(user *domain.User, role string) error
}
type Services struct {
	User  User
	Admin Admin
}
type Deps struct {
	Repos *repository.Repositories
}

//TODO 39.47
func NewServices(deps Deps) *Services {
	return &Services{User: NewUsersService(deps.Repos.Users), Admin: NewAdminsService(deps.Repos.Admins)}
}
