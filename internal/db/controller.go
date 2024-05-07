package db

import (
	"fmt"

	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseConfig struct {
	Host     string
	User     string
	Name     string
	Password string
	Port     string
}

func InitializeDbConnection(conf DatabaseConfig, logger *zap.Logger) *gorm.DB {
	connectionString := fmt.Sprintf("host=%s user=%s dbname=%s password=%s port=%s", conf.Host, conf.User, conf.Name, conf.Password, conf.Port)
	db, err := gorm.Open(
		postgres.Open(connectionString),
	)
	if err != nil {
		logger.Fatal("Could not initialize database",
			zap.Error(err),
		)
	}
	return db
}
