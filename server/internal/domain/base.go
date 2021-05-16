package domain

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	_ "gorm.io/gorm"
	"os"
)

func NewDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=require",
		"ec2-54-216-48-43.eu-west-1.compute.amazonaws.com",
		os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_NAME"),
		5432,
	)
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
