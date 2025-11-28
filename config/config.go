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
	Redis    RedisConfig
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

type RedisConfig struct {
	RedisHost     string 
    RedisPort     string 
    RedisPassword string 
    RedisDB       int    
}

func LoadConfig() *Config {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	// if err != nil {
	// 	log.Fatalf("Error reading config file: %v", err)
	// }

	if err != nil {
		// Jangan pakai Fatalf (Mati), pakai Println (Info) saja.
		// Karena di Docker, wajar kalau file .env tidak ditemukan.
		log.Println("⚠️  Warning: File .env tidak ditemukan. Menggunakan environment variables sistem.")
	} else {
		log.Println("✅  Sukses membaca konfigurasi dari file .env")
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
		Redis: RedisConfig{
			RedisHost:     viper.GetString("REDIS_HOST"),
			RedisPort:     viper.GetString("REDIS_PORT"),
			RedisPassword: viper.GetString("REDIS_PASSWORD"),
			RedisDB:       viper.GetInt("REDIS_DB"),
	
		},
	}
	return config
}