package user

import (
	"github.com/artemmarkaryan/gocrm/cmd/app/internal/domain"
	"github.com/artemmarkaryan/gocrm/cmd/app/pkg/jsons"
)

type UserPreview struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
}

func CreateUserPreview(user domain.User) *UserPreview {
	return &UserPreview{
		ID:       user.ID,
		Name:     user.Name,
		Username: user.Login,
	}
}

func (u UserPreview) Serialize() (string, error) {
	return jsons.MarshalToString(u)
}
