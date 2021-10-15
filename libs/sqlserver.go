package libs

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func (c *DbConfig) InitSqlServerDB() *gorm.DB {
	connString := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s", c.User, c.Password, c.Host, c.Port, c.Database)
	db, err := gorm.Open(sqlserver.Open(connString), &gorm.Config{})

	if err != nil {
		log.Panic(err)
		os.Exit(-1)
	}

	log.Println("SQL Server DB connected")
	return db
}
