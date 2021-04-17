package repository

import (
	"context"

	"github.com/IDarar/hub/internal/config"
	"github.com/IDarar/hub/internal/domain"
	"github.com/IDarar/hub/internal/repository/postgres"
	"github.com/IDarar/hub/internal/repository/redisdb"
	"github.com/go-redis/redis/v8"

	"gorm.io/gorm"
)

//all db interfaces there are described 44.20
//and struct repositories named with all interfaces in the end defined
type Users interface {
	Create(ctx context.Context, user domain.User) error
	GetRoleByID(id int) ([]string, error)
	GetByCredentials(ctx context.Context, name, password string) (domain.User, error)
	SetSession(userId int, session domain.Session) error

	AddTreatise(tr domain.UserTreatise) error
	UpdateTreatise(tr domain.UserTreatise) error
	AddProposition(pr domain.UserProposition) error
	UpdateProposition(pr domain.UserProposition) error
	AddPart(part domain.UserPart) error
	UpdatePart(part domain.UserPart) error
	RateTreatise(tr domain.UserTreatise, rate domain.Rate) error
	RatePart(part domain.UserPart, rate domain.Rate) error
	RateProposition(pr domain.UserProposition, rate domain.Rate) error
}
type Admins interface {
	GrantRole(name, role string) error

	RevokeRole(id int)
}
type Content interface {
	Create(treatise domain.Treatise) error
	Update(treatise domain.Treatise) error
	Delete(treatise domain.Treatise) error
	GetByID(id string) (domain.Treatise, error)
}
type Parts interface {
	Create(part domain.Part) error
	Update(part domain.Part, createLiterature, deleteLiterature []string) error
	Delete(part domain.Part) error
}
type Propositions interface {
	Create(proposition domain.Proposition) error
	Update(proposition domain.Proposition, createReferences, deleteReferences []string, createNotes, deleteNotes []domain.Note) error
	Delete(proposition domain.Proposition) error
	GetByID(id string) (domain.Proposition, error)
}
type Sessions interface {
	SetSession(userId int, session domain.Session, revoketoken string) error
	GetIDByRefreshToken(refreshToken string) (int, error)
	GetAllUserSessions(uID int) ([]string, error)
}
type Repositories struct {
	Users        Users
	Admins       Admins
	Content      Content
	Parts        Parts
	Propositions Propositions
	Sessions     Sessions
}

func NewRepositories(db *gorm.DB, rdb *redis.Client, cfg *config.Config) *Repositories {
	return &Repositories{
		Users:        postgres.NewUserRepo(db),
		Admins:       postgres.NewAdminsRepo(db),
		Content:      postgres.NewContentRepo(db),
		Parts:        postgres.NewPartsRepo(db),
		Propositions: postgres.NewPropositionsRepo(db),
		Sessions:     redisdb.NewSessionRepo(rdb, cfg),
	}
}
