package order

import (
	"errors"
	"github.com/artemmarkaryan/gocrm/cmd/app/internal/domain"
	"github.com/artemmarkaryan/gocrm/cmd/app/internal/dto/order"
	"github.com/artemmarkaryan/gocrm/cmd/app/internal/server/middleware"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
)

func (s Service) PatchOne(ctx *gin.Context) (err error) {
	var operation *gorm.DB

	db, err := domain.GetDB()
	if err != nil {
		return
	}

	session := sessions.Default(ctx)
	userIdFromSession, ok := session.Get(middleware.UserKey).(uint)
	if !ok {
		return errors.New("нет userId в сессии")
	}

	dbOrder := &domain.Order{UserID: userIdFromSession}
	operation = db.Find(&dbOrder, ctx.Param("id"))
	if operation.RowsAffected == 0 {
		return operation.Error
	}

	dtOrderPreview := *order.CreateOrderPreview(dbOrder)
	err = ctx.BindJSON(&dtOrderPreview)
	if err != nil {
		log.Fatal(err.Error())
	}

	dbOrder = dtOrderPreview.ToDBO()
	operation = db.Model(&dbOrder).Updates(dbOrder)
	err = operation.Error
	return
}
