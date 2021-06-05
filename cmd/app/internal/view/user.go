package view

import (
	"fmt"
	"github.com/artemmarkaryan/gocrm/cmd/app/internal/service/user"
	"net/http"
)

type UserView struct{}

func (u UserView) GetAll(
	w http.ResponseWriter,
	_ *http.Request,
) {
	users, _ := user.Service{}.GetAll()
	_, _ = fmt.Fprint(w, string(users))
}
