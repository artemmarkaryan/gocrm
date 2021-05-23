package main

import (
	"github.com/artemmarkaryan/gocrm/internal"
	"github.com/artemmarkaryan/gocrm/internal/domain"
	"log"
)

func main() {
	if err := internal.Setup(); err != nil {
		log.Panic("Configs not loaded:", err.Error())
	}
	db, err := domain.GetDB()
	if err != nil {
		log.Panic(err.Error())
	}
	_ = domain.Migrate(db)
	//result, err := service.UserService{}.GetAll()
	//if err != nil {
	//	log.Fatal(err.Error())
	//} else {
	//	log.Print(string(result))
	//}
}
