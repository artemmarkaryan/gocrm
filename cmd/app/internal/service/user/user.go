package user

import (
	"github.com/artemmarkaryan/gocrm/cmd/app/internal/domain"
	userDTO "github.com/artemmarkaryan/gocrm/cmd/app/internal/dto/user"
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
			*userDTO.NewUserPreview(u),
		)
	}

	return allUsersPreview.Serialize()
}
