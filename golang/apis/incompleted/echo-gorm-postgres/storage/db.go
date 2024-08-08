package storage

import (
	"log"

	"github.com/FelipeMatthew/go-learnings/echo-gorm/config"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func NewDB(parms...string) *gorm.DB {
	var err error
	conString := config.GetPostgresConnectionString()

	log.Print(conString)

	DB, err := gorm.Open(config.GetDbType(), conString)
	if err != nil {
		log.Panic(err)
	} 

	return DB
}

func GetDbInstance() *gorm.DB {
	return DB
}