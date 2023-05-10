package main

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"server/internal/local"
	"server/internal/repository"
	"server/internal/service"
	"server/internal/transport"
)

func main() {
	// Логирование и viper
	logrus.SetFormatter(new(logrus.JSONFormatter))
	initViperConfig()

	srv := new(local.Server)

	//handlers := new(transport.Handler)
	// Инциализируем подклбчение к Google Oauth2 API
	service.InitializeOAuthGoogle()

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		Password: viper.GetString("db.password"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		logrus.Fatal(err)
	}

	repos := repository.New(db)
	services := service.NewService(repos)
	handlers := transport.NewHandler(services)

	if err = srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}
}

func initViperConfig() {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	// Enable VIPER to read Environment Variables 😱🥹🔥
	viper.AutomaticEnv()

	viper.SetConfigType("yml")
	if err := viper.ReadInConfig(); err != nil {
		logrus.Fatalf("Error reading config file, %s", err)
	}
}
