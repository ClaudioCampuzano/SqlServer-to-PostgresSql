package libs

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func (c *DbConfig) InitPostgresDB() *gorm.DB {
	connString := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable", c.User, c.Password, c.Host, c.Port, c.Database)
	db, err := gorm.Open(postgres.Open(connString), &gorm.Config{})

	if err != nil {
		log.Panic(err)
		os.Exit(-1)
	}

	log.Println("Postgres DB connected")
	return db
}