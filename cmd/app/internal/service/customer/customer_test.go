package customer

import (
	"github.com/artemmarkaryan/gocrm/cmd/app/internal"
	"github.com/artemmarkaryan/gocrm/cmd/app/internal/domain"
	"github.com/artemmarkaryan/gocrm/cmd/app/pkg/jsonb"
	"strconv"
	"testing"
	"time"
)

func TestCustomerService_GetAll(t *testing.T) {
	if err := internal.Setup(); err != nil {
		t.Fatal(err.Error())
	}

	db, err := domain.GetDB()
	if err != nil {
		t.Fatal(err.Error())
	}
	testCustomer :=	 domain.Customer{
		Name: "test" + strconv.FormatInt(time.Now().Unix(), 10),
		Contact: jsonb.JSONB{"phone": "89991002030"},
	}
	db.Create(&testCustomer).Commit()

	service := CustomerService{}
	result, err := service.GetAll()
	if err != nil {
		t.Error(err.Error())
	} else {
		t.Log(string(result))
	}

	db.Delete(&testCustomer).Commit()
}
