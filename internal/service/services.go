package service

import (
	"context"
	"time"

	"github.com/IDarar/hub/internal/domain"
	"github.com/IDarar/hub/internal/repository"
	"github.com/IDarar/hub/pkg/auth"
	"github.com/IDarar/hub/pkg/hash"
)

type SignUpInput struct {
	Name     string
	Email    string
	Password string
}
type SignInInput struct {
	Name     string
	Password string
}

type Tokens struct {
	AccessToken  string
	RefreshToken string
}

//all interfaces there are described
type User interface {
	SignUp(ctx context.Context, input SignUpInput) error
	SignIn(ctx context.Context, input SignInInput) (Tokens, error)
	GetRoleById(id int) ([]string, error)
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
	Repos                  *repository.Repositories
	Hasher                 hash.PasswordHasher
	TokenManager           auth.TokenManager
	AccessTokenTTL         time.Duration
	RefreshTokenTTL        time.Duration
	CacheTTL               int64
	VerificationCodeLength int
}

//TODO 39.47
func NewServices(deps Deps) *Services {
	userService := NewUsersService(deps.Repos.Users, deps.Hasher,
		deps.TokenManager, deps.AccessTokenTTL, deps.RefreshTokenTTL, deps.VerificationCodeLength)
	return &Services{User: userService, Admin: NewAdminsService(deps.Repos.Admins)}
}
