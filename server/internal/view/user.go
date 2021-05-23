package view

import "github.com/artemmarkaryan/gocrm/internal/dto"

type UserView struct {}

func (u *UserView) GetAll() (result []byte, err error) {
	up := dto.AllUsersPreview{}
	up.Load();
	return up.Serialize()
}
