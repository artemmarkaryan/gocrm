package user

import (
	"encoding/json"
	"github.com/artemmarkaryan/gocrm/internal/domain"
)

type AllUsersPreview struct {
	UsersPreview []UserPreview
}

func (r AllUsersPreview) Serialize(
	users []domain.User,
) (result []byte, err error, ) {
	for _, user := range users {
		r.UsersPreview = append(
			r.UsersPreview,
			UserPreview{
				ID:   user.ID,
				Name: user.Name,
			})
	}
	return json.Marshal(r.UsersPreview)
}
