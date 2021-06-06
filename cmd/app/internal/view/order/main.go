package order

import (
	"fmt"
	"github.com/artemmarkaryan/gocrm/cmd/app/internal/service/order"
	"net/http"
)

type View struct {}

func (view View) GetAll(
	w http.ResponseWriter,
	_ *http.Request,
) {
	v, _ := order.Service{}.GetAll()
	_, _ = fmt.Fprint(w, v)
}
