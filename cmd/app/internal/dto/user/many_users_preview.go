package user

import (
	"github.com/artemmarkaryan/gocrm/cmd/app/pkg/jsons"
)

type ManyUsersPreview struct {
	UsersPreview []UserPreview
}

func (r ManyUsersPreview) Serialize() (string, error) {
	return jsons.MarshallToString(r.UsersPreview)
}
