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

	if err != nil {
		logger.Error(err)
		return

	}
	hasher := hash.NewSHA1Hasher(cfg.Auth.PasswordSalt)

	repos := repository.NewRepositories(db)

	//	createReferences := []string{"Third", "SECOND"}
	deleteReferences := []string{"EGQ", "LPE"}
	/*if len(createReferences) != 0 {

		for _, v := range createReferences {
			ref := domain.Reference{}
			err = db.Where(&domain.Reference{Target: "ACPWT", TargetProposition: v}).First(&ref).Error
			if err == nil {
				logger.Error("ref already exists")
				return //erros.New("ref already exists")
			}
			ref.Target = "ACPWT"
			ref.TargetProposition = v

			err = db.Create(&ref).Error
			if err != nil {
				logger.Error(err)
				return
			}
		}
	}*/
	if len(deleteReferences) != 0 {

		for _, v := range deleteReferences {
			ref := domain.Reference{}
			err = db.Where(&domain.Reference{Target: "EXVI", TargetProposition: v}).First(&ref).Error
			if err != nil {
				logger.Error("ref don't exist")
				return //erros.New("ref don't exist")
			}
			ref.Target = "EXVI"
			ref.TargetProposition = v

			count := db.Delete(&ref).RowsAffected
			if count == 0 {
				logger.Error("didn't deleted")
				return
			}
		}
	}
	services := service.NewServices(service.Deps{
		Repos:           repos,
		Hasher:          hasher,
		AccessTokenTTL:  cfg.Auth.JWT.AccessTokenTTL,
		RefreshTokenTTL: cfg.Auth.JWT.RefreshTokenTTL,
		TokenManager:    tokenManager,
	})

	handlers := http.NewHandler(services, tokenManager)
	srv := server.NewServer(cfg, handlers.Init())

	srv.Run()
}

/*
pr, err := repos.Propositions.GetByID("ACPWT")
	if err != nil {
		logger.Error(err)
		return

	}
	res := strings.Contains(pr.Text, "xxvi")
	fmt.Println(res) // true

	i := strings.Index(pr.Text, "xxvi")
	fmt.Println(i)

tr, err := repos.Content.GetByID("ERRRR")
if err != nil {
logger.Error(err)
return

}
logger.Info(tr)

err := r.db.Save(&prop).Error
if err != nil {
logger.Error(err)
return err

}

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
