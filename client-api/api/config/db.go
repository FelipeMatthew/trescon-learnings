package config

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	user     = "root"
	password = "password"
	dbName   = "client_api"
	port     = "5432"

	database *gorm.DB
	e        error
)

func DatabaseInit() {

	dns := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbName, port)

	database, e = gorm.Open(postgres.Open(dns), &gorm.Config{})
	if e != nil {
		panic(e)
	}

	log.Println(dns)
}

func DB() *gorm.DB {
	return database
}
