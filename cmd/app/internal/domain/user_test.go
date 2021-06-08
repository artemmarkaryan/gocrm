package domain

import (
	"github.com/artemmarkaryan/gocrm/cmd/app/internal"
	"testing"
	"time"
)

func TestUser(t *testing.T) {
	if err := internal.Setup(); err != nil {
		t.Fatal(err.Error())
	}

	db, err := GetDB()
	if err != nil {
		t.Fatal(err.Error())
		return
	}
	db.Create(&User{
		Name: time.Now().String(),
	}).Commit()
}