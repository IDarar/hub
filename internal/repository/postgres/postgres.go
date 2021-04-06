package postgres

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(config Config) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open("" /*TODO make one string with configs*/), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return db, nil
}
