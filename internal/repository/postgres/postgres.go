package postgres

import (
	"fmt"

	"github.com/IDarar/hub/internal/config"
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

/*type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string

	user=postgres dbname=forum password=123 sslmode=disabled

	SSLMode  string
}*/

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
	err := db.AutoMigrate(&User{}, &Session{})
	if err != nil {
		return
	}
	logger.Info("migrated succsessfully")

}
