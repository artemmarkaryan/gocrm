package main

import (
	"github.com/artemmarkaryan/gocrm/cmd/app/internal"
	"github.com/artemmarkaryan/gocrm/cmd/app/internal/domain"
	"log"
)

func main() {
	if err := internal.Setup(); err != nil {
		log.Print("Configs not loaded:", err.Error())
	}
	db, err := domain.GetDB()
	if err != nil {
		log.Panic(err.Error())
	}
	err = domain.Migrate(db)
	if err != nil {
		log.Panic(err.Error())
	}

	//server.Run()
}
