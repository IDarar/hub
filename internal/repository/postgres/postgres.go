package postgres

import (
	"fmt"

	"github.com/IDarar/hub/internal/config"
	"github.com/IDarar/hub/internal/domain"
	"github.com/IDarar/hub/pkg/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresConfig struct {
	DBname   string
	User     string
	Password string
	Sslmode  string
}

func NewPostgresDB(cfg *config.Config) (*gorm.DB, error) {
	fmt.Println("user=" + cfg.Postgres.User + " " + "dbname=" + cfg.Postgres.DBname + " " + "password=" + cfg.Postgres.Password + " " + "sslmode=" + cfg.Postgres.Sslmode)
	db, err := gorm.Open(postgres.Open("user="+cfg.Postgres.User+" "+"dbname="+cfg.Postgres.DBname+" "+"password="+cfg.Postgres.Password+" "+"sslmode="+cfg.Postgres.Sslmode), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println(cfg.Postgres)
	InitialiseTables(db)
	return db, nil
}
func InitialiseTables(db *gorm.DB) {
	err := db.AutoMigrate(&domain.User{},
		&domain.Session{},
		&domain.Role{},
		&domain.Treatise{},
		&domain.Part{},
		&domain.Proposition{},
		&domain.Reference{},
		&domain.Literature{},
		&domain.Note{},
		&domain.Rate{},
		domain.UserLists{},
		domain.UserTreatise{},
		domain.UserPart{},
		domain.UserProposition{},
		domain.UserNote{})
	if err != nil {
		return
	}

	logger.Info("migrated succsessfully")

}

func AssociationCreate(db *gorm.DB, str interface{}, association string, toAppend interface{}) error {
	err := db.Model(&str).Association("Propositions").Append(&toAppend)
	if err != nil {
		logger.Error(err)
		return err
	}

	return nil
}
func CheckExByID(db *gorm.DB, toCheck interface{}, condID ...interface{}) error {
	err := db.Model(toCheck).First(&toCheck).Error
	if err != nil {
		logger.Error(err)
		return err
	}
	return nil
}
