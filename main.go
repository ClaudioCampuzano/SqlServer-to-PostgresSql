package main

import (
	"github.com/ClaudioCampuzano/Golang-SqlServer-postgreSQL/libs"
)

func main() {
	dbConfig := libs.Configure("./", "postgres")
	libs.DB = dbConfig.InitPostgresDB()
}
