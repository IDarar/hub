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
type AddTreatiseInput struct {
	TargetTreatise string
}
type UpdateUserTreatise struct {
	TargetTreatise string
	Status         string
	IsCompleted    *bool
}
type AddPropositionInput struct {
	TargetProposition string
}
type UpdateUserProposition struct {
	TargetProposition string
	Status            string
	IsCompleted       *bool
}
type AddPartInput struct {
	TargetPart string
}
type UpdateUserPart struct {
	TargetPart  string
	Status      string
	IsCompleted *bool
}
type RateInput struct {
	Target string
	Type   string
	Value  int
}
type User interface {
	SignUp(ctx context.Context, input SignUpInput) error
	SignIn(ctx context.Context, input SignInInput) (Tokens, error)
	RefreshTokens(refreshToken string) (Tokens, error)
	GetRoleById(id int) ([]string, error)
	CreateMark(domain.UserProposition, [3]interface{}) error
	AddTreatise(inp AddTreatiseInput, userID interface{}) error
	UpdateTreatise(inp UpdateUserTreatise, userID interface{}) error
	AddProposition(inp AddPropositionInput, userID interface{}) error
	UpdateProposition(inp UpdateUserProposition, userID interface{}) error
	AddPart(inp AddPartInput, userID interface{}) error
	UpdatePart(inp UpdateUserPart, userID interface{}) error
	RateTreatise(rate RateInput, userID interface{}) error
	RatePart(rate RateInput, userID interface{}) error
	RateProposition(rate RateInput, userID interface{}) error
	DeleteRateTreatise(rate RateInput, userID interface{}) error
	DeleteRatePart(rate RateInput, userID interface{}) error
	DeleteRateProposition(rate RateInput, userID interface{}) error
}

type RoleInput struct {
	UserName string
	Role     string
}
type Admin interface {
	GrantRole(name, role string, roles interface{}) error
	RevokeRole(user *domain.User, role string) error
}

type TreatiseUpdateInput struct {
	ID          string
	Title       string
	Description string
	Date        string
}
type Content interface {
	Create(id, title, date, description string, roles interface{}) error
	Delete(id string, roles interface{}) error
	Update(inp TreatiseUpdateInput, roles interface{}) error
}
type PartUpdateInput struct {
	ID               string
	Name             string
	FullName         string
	TargetID         string
	Description      string
	CreateLiterature []string
	DeleteLiterature []string
}
type Part interface {
	Create(id, TargetID, name, fullname, description string, roles interface{}) error
	Delete(id string, roles interface{}) error
	Update(inp PartUpdateInput, roles interface{}) error
	AddToFavourite(fav domain.Favourite) error
	RemoveFromFavourite(fav domain.Favourite) error
}
type CreateProposition struct {
	ID          string
	TargetID    string
	Name        string
	Description string
	Explanation string
	Text        string
}
type UpdatePropositionInput struct {
	ID               string
	TargetID         string
	Name             string
	Description      string
	Explanation      string
	Text             string
	CreateReferences []string
	DeleteReferences []string
	CreateNotes      []domain.Note
	DeleteNotes      []domain.Note
}
type Propositions interface {
	Create(prop CreateProposition, roles interface{}) error
	Delete(id string, roles interface{}) error
	Update(inp UpdatePropositionInput, roles interface{}) error
	AddToFavourite(fav domain.Favourite) error
	RemoveFromFavourite(fav domain.Favourite) error
}

//Another way of implemntation
func Favourite(ID string, fav repository.Favourite) {
	fav.AddToFavourite(fav)
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
	userService := NewUsersService(deps.Repos.Users, deps.Repos.Sessions, deps.Hasher,
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
