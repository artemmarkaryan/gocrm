package order

import (
	"errors"
	"github.com/artemmarkaryan/gocrm/cmd/app/internal/domain"
	"github.com/artemmarkaryan/gocrm/cmd/app/internal/dto/order"
	"github.com/artemmarkaryan/gocrm/cmd/app/internal/server/middleware"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func (s Service) New(ctx *gin.Context) (orderId *uint, err error) {
	db, err := domain.GetDB()
	if err != nil {
		return
	}
	session := sessions.Default(ctx)
	userIdFromSession, ok := session.Get(middleware.UserKey).(uint)
	if !ok {
		err = errors.New("нет userId в сессии")
		return
	}

	dtOrder := order.NewOrder{UserId: userIdFromSession}
	err = ctx.BindJSON(&dtOrder)

	if err != nil {
		return
	}

	dbOrder := dtOrder.ToDBO()
	operation := db.Create(dbOrder).Commit()
	if operation.RowsAffected == 0{
		err = operation.Error
	}
	return &dbOrder.ID, err
}

