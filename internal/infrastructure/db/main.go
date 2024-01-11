package db

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	var err error
	err = godotenv.Load()
	if err != nil {
		fmt.Println(".env file could not be loaded")
	}
	// dsn := "host=db user=admin password=admin dbname=tuix_client port=5432 sslmode=disable TimeZone=UTC"
	dsn := "host=" + os.Getenv("POSTGRES_HOST") +
		" user=" + os.Getenv("POSTGRES_USER") +
		" password=" + os.Getenv("POSTGRES_PASSWORD") +
		" dbname=" + os.Getenv("POSTGRES_DB") +
		" port=" + os.Getenv("POSTGRES_PORT") +
		" sslmode=" + os.Getenv("SSL_MODE") +
		" TimeZone=" + os.Getenv("TIMEZONE")
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("CONNECTED TO DB")

	// Migrate the schema
	db.AutoMigrate(&User{})

}

func DB() *gorm.DB {
	return db
}
