package user

import (
	"errors"
	"github.com/artemmarkaryan/gocrm/cmd/app/internal/domain"
	userDTO "github.com/artemmarkaryan/gocrm/cmd/app/internal/dto/user"
	"github.com/artemmarkaryan/gocrm/cmd/app/pkg/encryption"
	"os"
)

var (
	EWrongPassword = errors.New("Wrong password")
)

type Service struct{}

func (userService Service) GetAll() (result string, err error) {
	db, err := domain.GetDB()
	if err != nil {
		return
	}

	var users []domain.User
	db.Find(&users)

	allUsersPreview := userDTO.ManyUsersPreview{}
	for _, u := range users {
		allUsersPreview.UsersPreview = append(
			allUsersPreview.UsersPreview,
			*userDTO.CreateUserPreview(u),
		)
	}

	return allUsersPreview.Serialize()
}

func (userService Service) Auth(username, password string) (ok bool, err error) {
	db, err := domain.GetDB()
	if err != nil {
		return
	}

	var user domain.User
	encryptedPassword := encryption.Encrypt(password, os.Getenv("SECRET"))
	db.Find(&user, "login = ? and password = ?", username, encryptedPassword)
	if user.ID == 0 {
		return false, EWrongPassword
	} else {
		return true, nil
	}
}