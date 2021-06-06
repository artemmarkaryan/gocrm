package user

import (
	"fmt"
	"github.com/artemmarkaryan/gocrm/cmd/app/internal/service/user"
	"net/http"
)

type View struct{}

func (view View) GetAll(
	w http.ResponseWriter,
	_ *http.Request,
) {
	v, _ := user.Service{}.GetAll()
	_, _ = fmt.Fprint(w, v)
}
