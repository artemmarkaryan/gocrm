package user

import (
	"github.com/artemmarkaryan/gocrm/src/internal/domain"
)

type UserPreview struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func (u *UserPreview) Serialize(user ...domain.DBObject) (result []byte, err error) {
	panic("implement me")
}
