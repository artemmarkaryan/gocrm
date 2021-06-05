package user

import (
	"github.com/artemmarkaryan/gocrm/cmd/app/internal"
	"testing"
)

func TestUserService_GetAll(t *testing.T) {
	if err := internal.Setup(); err != nil {
		t.Fatal(err.Error())
	}

	service := Service{}
	result, err := service.GetAll()
	if err != nil {
		t.Error(err.Error())
	} else {
		t.Log(string(result))
	}

}
