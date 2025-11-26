package databases

import (
	"fmt"
	"log"
	"tokogue-api/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDBConnection(config *config.Config) *gorm.DB {
dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Jakarta",
	config.Databases.Host,
	config.Databases.User,
	config.Databases.Password,
	config.Databases.DBName,
	config.Databases.Port,
	config.Databases.SSLMode,
)
db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
if err != nil {
	log.Fatalf("Failed to connect to database: %v", err)
}
return db
}