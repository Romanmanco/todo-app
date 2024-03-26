package main

import (
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
	"os"
	"todo-app/pkg/handler"
	"todo-app/pkg/server"
	"todo-app/pkg/service"
	"todo-app/pkg/storage"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error config initializing: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error getting .env var: %s", err.Error())
	}

	db, err := storage.NewPostgresDB(storage.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		log.Fatalf("failed initializing db: %s", err.Error())
	}

	repo := storage.NewRepository(db)
	services := service.NewService(repo)
	handlers := handler.NewHandler(services)

	serv := new(server.Server)
	if err := serv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error while run http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
