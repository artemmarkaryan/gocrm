package server

import (
	"github.com/artemmarkaryan/gocrm/cmd/app/internal/view/customer"
	"github.com/artemmarkaryan/gocrm/cmd/app/internal/view/order"
	"github.com/artemmarkaryan/gocrm/cmd/app/internal/view/product"
	"github.com/artemmarkaryan/gocrm/cmd/app/internal/view/user"
	"log"
	"net/http"
)

const address = "0.0.0.0:8000"

func Run() {
	http.HandleFunc("/user/all", user.View{}.GetAll)
	http.HandleFunc("/customer/all", customer.View{}.GetAll)
	http.HandleFunc("/product/all", product.View{}.GetAll)
	http.HandleFunc("/order/all", order.View{}.GetAll)

	log.Printf("Running on http://%v", address)
	_ = http.ListenAndServe(address, nil)
}
