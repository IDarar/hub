package repository

import (
	"context"

	"github.com/IDarar/hub/internal/domain"
	"github.com/IDarar/hub/internal/repository/postgres"

	"gorm.io/gorm"
)

//all db interfaces there are described 44.20
//and struct repositories named with all interfaces in the end defined
type Users interface {
	Create(ctx context.Context, user domain.User) error
	GetRoleByID(id int) ([]string, error)
	GetByCredentials(ctx context.Context, name, password string) (domain.User, error)
	SetSession(userId int, session domain.Session) error

	//TODO CreateMark(map[int]int)
}
type Admins interface {
	GrantRole(name, role string) error

	RevokeRole(id int)
}
type Content interface {
	Create(treatise domain.Treatise) error
	Delete(treatise domain.Treatise) error
}
type Repositories struct {
	Users   Users
	Admins  Admins
	Content Content
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		Users:   postgres.NewUserRepo(db),
		Admins:  postgres.NewAdminsRepo(db),
		Content: postgres.NewContentRepo(db),
	}
}
