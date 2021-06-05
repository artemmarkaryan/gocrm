package view

import (
	"github.com/artemmarkaryan/gocrm/src/internal/service"
)

type UserView struct{}

func (u *UserView) GetAll() (result []byte, err error) {
	return service.UserService{}.GetAll()
}
