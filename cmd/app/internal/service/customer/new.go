package customer

import (
	"github.com/artemmarkaryan/gocrm/cmd/app/internal/domain"
	"github.com/artemmarkaryan/gocrm/cmd/app/internal/dto/customer"
	"github.com/gin-gonic/gin"
)

func (r Service) New(ctx *gin.Context) (customerId uint, err error) {
	var customerDBO domain.Customer
	var customerDTO customer.Customer

	err = ctx.BindJSON(&customerDTO)
	if err != nil {
		return 0, err
	}

	customerDBO = *customerDTO.ToDBO()

	db, err := domain.GetDB()
	if err != nil {
		return
	}

	operation := db.Create(&customerDBO).Commit()
	if operation.RowsAffected == 0 {
		err = operation.Error
	}

	return customerDBO.ID, err
}
