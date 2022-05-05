package database

import (
	"github.com/yamess/inventory/configs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type myDatabase struct {
	Conn *gorm.DB
}

var MyDB = myDatabase{}

func (pgDB *myDatabase) Connect() {
	db, err := gorm.Open(postgres.Open(configs.PgDbUrl), &gorm.Config{})

	if err != nil {
		log.Println(err.Error())
		panic("Failed to connect to database")
	}

	pgDB.Conn = db

}

func Automigrate(models ...interface{}) {
	MyDB.Connect()
	for _, v := range models {
		ok := MyDB.Conn.AutoMigrate(&v)
		if ok != nil {
			log.Println(ok.Error())
			panic("Failed to apply migration")
		}
	}
}
