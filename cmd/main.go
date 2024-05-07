package main

import (
	"log"
	"net"

	"github.com/pstano1/customer-api/internal/db"
	"github.com/pstano1/customer-api/service/api"
	handler "github.com/pstano1/customer-api/service/grpc"
	"github.com/pstano1/customer-api/service/repository"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

const (
	confOptBindPort         = "BIND_PORT"
	confOptDatabaseHost     = "DATABASE_HOST"
	confOptDatabaseName     = "DATABASE_NAME"
	confOptDatabaseUsername = "DATABASE_USERNAME"
	confOptDatabasePassword = "DATABASE_PASSWORD"
	confOptDatabasePort     = "DATABASE_PORT"
)

func main() {
	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	logger := zap.Must(zap.NewProduction())
	defer logger.Sync()
	conf := &db.DatabaseConfig{
		User:     viper.GetString(confOptDatabaseUsername),
		Password: viper.GetString(confOptDatabasePassword),
		Host:     viper.GetString(confOptDatabaseHost),
		Port:     viper.GetString(confOptDatabasePort),
		Name:     viper.GetString(confOptDatabaseName),
	}
	database := db.InitializeDbConnection(*conf, logger)

	lis, err := net.Listen("tcp", viper.GetString(confOptBindPort))
	if err != nil {
		logger.Fatal("error while starting the server",
			zap.Error(err),
		)
	}

	grpcServer := grpc.NewServer()

	customersRepository := repository.New(database)
	API := api.New(customersRepository)
	handler.NewService(grpcServer, API)

	log.Fatal(grpcServer.Serve(lis))
}
