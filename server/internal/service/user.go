package service

import (
	"github.com/artemmarkaryan/gocrm/internal/domain"
	userDTO "github.com/artemmarkaryan/gocrm/internal/dto/user"
)

type UserService struct {}

func (userService UserService) GetAll() (result []byte, err error) {
	db, err := domain.GetDB()
	if err != nil {
		return
	}

	var users []domain.User
	db.Find(&users)
	return userDTO.AllUsersPreview{}.Serialize(users)
}

