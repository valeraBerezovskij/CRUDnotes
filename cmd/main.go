package main

import (
	"os"
	"valeraninja/noteapp/internal/handlers"
	"valeraninja/noteapp/internal/repository"
	"valeraninja/noteapp/internal/server"
	"valeraninja/noteapp/internal/services"
	"valeraninja/noteapp/pkg/database"

	"github.com/sirupsen/logrus"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

// @title CRUD APP
// @version 1.1
// @description CRUD application for notes

// @host localhost:8000
// @BasePath /
func main() {
	initLogger()

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env file: %s", err.Error())
	}

	db, err := database.NewPostgresDB(database.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USERNAME"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		Password: os.Getenv("DB_PASSWORD"),
	})

	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := services.NewService(repos)
	handlers := handlers.NewHandler(services.NoteItem)

	srv := new(server.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initLogger() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.InfoLevel)
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
