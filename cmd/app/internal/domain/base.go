package domain

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"sync"
)

var once sync.Once
var DB *gorm.DB

func GetDB() (*gorm.DB, error) {
	var err error
	once.Do(func() {
		dsn := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%d sslmode=require",
			"ec2-54-216-48-43.eu-west-1.compute.amazonaws.com",
			os.Getenv("DATABASE_USER"),
			os.Getenv("DATABASE_PASSWORD"),
			os.Getenv("DATABASE_NAME"),
			5432,
		)
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	})
	return DB, err
}

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&Customer{},
		&Product{},
		&Basket{},
		&Item{},
		&User{},
		&Order{},
	)
}
