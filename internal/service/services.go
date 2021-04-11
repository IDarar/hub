package service

import (
	"context"
	"time"

	_ "github.com/golang/mock/mockgen/model"

	"github.com/IDarar/hub/internal/domain"
	"github.com/IDarar/hub/internal/repository"
	"github.com/IDarar/hub/pkg/auth"
	"github.com/IDarar/hub/pkg/hash"
)

//go:generate go run -mod=mod github.com/golang/mock/mockgen -package mocks -destination=mocks/mock.go -source=services.go -build_flags=-mod=mod

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

type RoleInput struct {
	UserName string
	Role     string
}
type Admin interface {
	GrantRole(name, role string, roles interface{}) error
	RevokeRole(user *domain.User, role string) error
}

type Content interface {
	Create(id, title, date, description string, roles interface{}) error
	Delete(id string, roles interface{}) error
}

type Part interface {
	Create(id, TargetID, name, fullname, description string, roles interface{}) error
	Delete(id string, roles interface{}) error
}
type CreateProposition struct {
	ID          string
	TargetID    string
	Name        string
	Description string
	Explanation string
	Text        string
}
type Propositions interface {
	Create(prop CreateProposition, roles interface{}) error
	Delete(id string, roles interface{}) error
}

type Services struct {
	User         User
	Admin        Admin
	Content      Content
	Part         Part
	Propositions Propositions
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
	adminService := NewAdminsService(deps.Repos.Admins, userService)
	contentService := NewContentService(deps.Repos.Content, userService)
	partsService := NewPartsService(deps.Repos.Parts, userService)
	propositionsService := NewPropositionsService(deps.Repos.Propositions, userService)
	return &Services{
		User:         userService,
		Admin:        adminService,
		Content:      contentService,
		Part:         partsService,
		Propositions: propositionsService}
}
