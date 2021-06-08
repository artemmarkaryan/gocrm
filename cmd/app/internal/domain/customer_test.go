package domain

import (
	"github.com/artemmarkaryan/gocrm/cmd/app/internal"
	"github.com/artemmarkaryan/gocrm/cmd/app/pkg/jsonb"
	"testing"
	"time"
)

func TestCustomer(t *testing.T) {
	if err := internal.Setup(); err != nil {
		t.Fatal(err.Error())
	}

	db, err := GetDB()
	if err != nil {
		t.Fatal(err.Error())
		return
	}
	db.Create(&Customer{
		Name: time.Now().String(),
		Contact: &jsonb.JSONB{
			"phone": "89150641853",
		},
	}).Commit()
}