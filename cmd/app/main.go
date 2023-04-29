package main

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"server/internal/handler"
	"server/internal/local"
	"server/internal/service"
)

func main() {
	// Логирование и viper
	logrus.SetFormatter(new(logrus.JSONFormatter))
	initViperConfig()

	srv := new(local.Server)

	handlers := new(handler.Handler)
	// Инциализируем подклбчение к Google Oauth2 API
	service.InitializeOAuthGoogle()

	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}
}

func initViperConfig() {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	viper.SetConfigType("yml")
	if err := viper.ReadInConfig(); err != nil {
		logrus.Fatalf("Error reading config file, %s", err)
	}
}
