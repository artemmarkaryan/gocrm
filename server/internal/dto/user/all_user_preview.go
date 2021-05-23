package user

import (
	"encoding/json"
	"github.com/artemmarkaryan/gocrm/internal/domain"
)

type AllUsersPreview struct {
	UsersPreview []UserPreview
}

func (AllUsersPreview AllUsersPreview) Serialize(
	users []domain.User,
) (result []byte, err error, ) {
	for _, user := range users {
		AllUsersPreview.UsersPreview = append(
			AllUsersPreview.UsersPreview,
			UserPreview{
				ID:   uint(user.ID),
				Name: user.Name,
			})
	}
	return json.Marshal(AllUsersPreview.UsersPreview)
}
