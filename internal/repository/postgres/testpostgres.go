package postgres

import (
	"fmt"
	"strings"
	"testing"

	"github.com/IDarar/hub/internal/config"
	"github.com/IDarar/hub/pkg/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const configPath = "configs/main"

func NewTestPostgres(t *testing.T) (*gorm.DB, func(...string)) {
	cfg, err := config.Init(configPath)
	logger.Info("CONFIG", cfg)
	if err != nil {
		logger.Error(err)
		return nil, nil
	}
	db, err := NewTestPostgresDB(cfg)
	if err != nil {
		logger.Error(err)
		return nil, nil
	}
	InitialiseTables(db)

	return db, func(tables ...string) {
		if len(tables) > 0 {
			if err := db.Exec(fmt.Sprintf("TRUNCATE %s CASCADE", strings.Join(tables, ", "))).Error; err != nil {
				t.Fatal(err)
			}
		}

	}

}
func NewTestPostgresDB(cfg *config.Config) (*gorm.DB, error) {
	logger.Info(cfg.TestPostgresConfig)
	db, err := gorm.Open(postgres.Open(cfg.TestPostgresConfig.URL), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println(cfg.Postgres)
	InitialiseTables(db)
	return db, nil
}
