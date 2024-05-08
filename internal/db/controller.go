package db

import (
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializeDbConnection(connectionString string, logger *zap.Logger) *gorm.DB {
	db, err := gorm.Open(
		postgres.Open(connectionString),
		&gorm.Config{},
	)
	if err != nil {
		logger.Fatal("Could not initialize database",
			zap.Error(err),
		)
	}
	return db
}
