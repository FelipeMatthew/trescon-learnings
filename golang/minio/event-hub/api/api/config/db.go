package config

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	database *gorm.DB
	e        error
)

func DatabaseInit() {
	host := "localhost"
	user := "postgres"
	password := "password"
	dbName := "eventhub"
	port := "5432"

	dns := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbName, port)

	database, e = gorm.Open(postgres.Open(dns), &gorm.Config{})
	if e != nil {
		panic(e)
	}

	_, err := database.DB()
	if err != nil {
		log.Fatal("failed to get database object from GORM DB:", err)
	}
}

func DB() *gorm.DB {
	return database
}
