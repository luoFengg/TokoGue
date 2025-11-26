package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Server ServerConfig
	Databases DatabasesConfig
	JWT      JWTConfig
	Midtrans MidtransConfig
}

type JWTConfig struct {
	Secret string
}

type DatabasesConfig struct {
	Host	 string
	Port	 string
	User	 string
	Password string
	DBName	 string
	SSLMode string
}


type ServerConfig struct {
	Host string
	Port string
}

type MidtransConfig struct {
	ServerKey string
	ClientKey string
}

func LoadConfig() *Config {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	config := &Config{
		Server: ServerConfig{
			Host: viper.GetString("HOST"),
			Port: viper.GetString("PORT"),
		},
		Databases: DatabasesConfig{
			Host:     viper.GetString("DB_HOST"),
			Port:     viper.GetString("DB_PORT"),
			User:     viper.GetString("DB_USER"),
			Password: viper.GetString("DB_PASSWORD"),
			DBName:   viper.GetString("DB_NAME"),
			SSLMode:  viper.GetString("DB_SSLMODE"),
		},
		JWT: JWTConfig{
			Secret: viper.GetString("JWT_SECRET"),
		},
		Midtrans: MidtransConfig{
			ServerKey: viper.GetString("MIDTRANS_SERVER_KEY"),
			ClientKey: viper.GetString("MIDTRANS_CLIENT_KEY"),
		},
	}
	return config
}