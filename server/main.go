package main

import (
	"github.com/artemmarkaryan/gocrm/internal/domain"
	dotenv "github.com/joho/godotenv"
	"log"
)

func configs() error {
	return dotenv.Load()
}


func main() {
	if err := configs(); err != nil {
		log.Panic("Configs not loaded:", err.Error())
	}
	_, _ = domain.NewDB()
}
