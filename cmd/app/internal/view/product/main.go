package product

import (
	"fmt"
	"github.com/artemmarkaryan/gocrm/cmd/app/internal/service/product"
	"net/http"
)

type View struct{}

func (view View) GetAll(
	w http.ResponseWriter,
	_ *http.Request,
) {
	v, _ := product.Service{}.GetAll()
	_, _ = fmt.Fprint(w, v)
}
