package app

import (
	"github.com/IDarar/hub/internal/config"
	"github.com/IDarar/hub/internal/domain"
	"github.com/IDarar/hub/internal/repository"
	"github.com/IDarar/hub/internal/repository/postgres"
	"github.com/IDarar/hub/internal/server"
	"github.com/IDarar/hub/internal/service"
	"github.com/IDarar/hub/internal/transport/http"
	"github.com/IDarar/hub/pkg/auth"
	"github.com/IDarar/hub/pkg/hash"
	"github.com/IDarar/hub/pkg/logger"
)

// @title Hub
// @version 0.001
// @description Hub for specified topics

// @host localhost:8080
// @BasePath /api/v1/

// @securityDefinitions.apikey AdminAuth
// @in header
// @name Authorization

// @securityDefinitions.apikey StudentsAuth
// @in header
// @name Authorization
func Run(configPath string) {
	cfg, err := config.Init(configPath)
	if err != nil {
		logger.Error(err)
		return
	}
	tokenManager, err := auth.NewManager(cfg.Auth.JWT.SigningKey)
	if err != nil {
		logger.Error(err)
		return
	}
	db, err := postgres.NewPostgresDB(cfg)
	if err != nil {
		logger.Error(err)
		return
	}
	hasher := hash.NewSHA1Hasher(cfg.Auth.PasswordSalt)

	repos := repository.NewRepositories(db)
	services := service.NewServices(service.Deps{
		Repos:           repos,
		Hasher:          hasher,
		AccessTokenTTL:  cfg.Auth.JWT.AccessTokenTTL,
		RefreshTokenTTL: cfg.Auth.JWT.RefreshTokenTTL,
		TokenManager:    tokenManager,
	})
	treatise := &domain.Treatise{ID: "E"}
	/*err = db.Create(&treatise).Error
	if err != nil {
		logger.Error(err)
		return

	}*/
	prop := &domain.Proposition{ID: "14EXVI", TargetID: "E3"}
	err = db.Create(&prop).Error
	if err != nil {
		logger.Error(err)
		return

	}
	err = db.Model(&treatise).Association("Propositions").Append(prop)
	if err != nil {
		logger.Error(err)
		return

	}
	err = db.Model(&prop).Association("Propositafafions").Append([]domain.Treatise{})
	if err != nil {
		logger.Error(err)
		return

	}
	handlers := http.NewHandler(services, tokenManager)
	srv := server.NewServer(cfg, handlers.Init())

	srv.Run()
}

/*
treatise := &domain.Treatise{ID: part.TargetID}
	toin := &domain.Part{Name: part.Name, ID: part.ID}
	err = db.Model(&treatise).Association("Parts").Append(&toin)
	if err != nil {
		logger.Error(err)
		return
	}
	//err := r.db.Model(&domain.Part{}).Where("target_id = ?", part.TargetID).Updates(domain.Part{ID: part.ID, Name: part.Name}).Error
role := domain.Role{Role: "SuperModerator"}
db.Create(&role)
role = domain.Role{Role: "SuperModerator", Users: []domain.User{{ID: 1}, {ID: 2}}}

db.Model(&role).Association("Users").Append([]domain.User{})

repos.Users.GetRoleByID(20)
repos.Users.GetRoleByID(3)

role := domain.Role{Role: "ContentModerator"}
role2 := domain.Role{Role: "ForumModerator"}

db.Create(&role)
db.Create(&role2)
*/
