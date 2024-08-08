package config

import (
	"fmt"
	"log"

	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)

var database *gorm.DB
var e error

func DatabaseInit() {
	host := "localhost"
	user := "root"
	password := "password"
	dbName := "library"
	port := "5432"

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