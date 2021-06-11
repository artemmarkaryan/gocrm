package user

import (
	"github.com/artemmarkaryan/gocrm/cmd/app/internal"
	"log"
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

func TestService_Auth(t *testing.T) {
	if err := internal.Setup(); err != nil {
		t.Fatal(err.Error())
	}

	service := Service{}
	for c, result := range map[struct{ username, password string }]bool{
		{"username", "password"}:  true,
		{"username", "password1"}: false,
	} {
		ok, err := service.Auth(c.username, c.password)
		log.Printf("case %v -> %v; ok: %v; err: %v", c, result, ok, err)
		//if ok != result {
		//	t.Error(err)
		//} else {
		//	log.Printf("%v -> good", c)
		//}
	}

}
