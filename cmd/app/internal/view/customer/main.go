package customer

import (
	"fmt"
	"github.com/artemmarkaryan/gocrm/cmd/app/internal/service/customer"
	"net/http"
)

type View struct{}

func (view View) GetAll(
	w http.ResponseWriter,
	_ *http.Request,
) {
	v, _ := customer.Service{}.GetAll()
	_, _ = fmt.Fprint(w, v)
}
