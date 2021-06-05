package view

import (
	"fmt"
	"github.com/artemmarkaryan/gocrm/cmd/app/internal/service/order"
	"net/http"
)

type OrderView struct {}

func (o OrderView) GetAll(
	w http.ResponseWriter,
	_ *http.Request,
) {
	orders, _ := order.Service{}.GetAll()
	_, _ = fmt.Fprint(w, string(orders))
}
