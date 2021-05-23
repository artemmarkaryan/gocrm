package main

import (
	"github.com/artemmarkaryan/gocrm/internal"
	"github.com/artemmarkaryan/gocrm/internal/service"
	"log"
)

func main() {
	if err := internal.Setup(); err != nil {
		log.Panic("Configs not loaded:", err.Error())
	}
	result, err := service.UserService{}.GetAll()
	if err != nil {
		log.Fatal(err.Error())
	} else {
		log.Print(string(result))
	}
}
