package server

import (
	"github.com/artemmarkaryan/gocrm/cmd/app/internal/view"
	"net/http"
)

func Run() {
	http.HandleFunc("/user/all", view.UserView{}.GetAll)

	_ = http.ListenAndServe("0.0.0.0:8000", nil)
}