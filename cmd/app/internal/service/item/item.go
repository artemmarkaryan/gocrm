package item

import (
	"errors"
	"fmt"
	"github.com/artemmarkaryan/gocrm/cmd/app/internal/domain"
	"github.com/artemmarkaryan/gocrm/cmd/app/internal/dto/item"
	"github.com/artemmarkaryan/gocrm/cmd/app/internal/service"
	"github.com/gin-gonic/gin"
)

type Service struct{}

func (s Service) PatchOne(ctx *gin.Context) (result string, err error) {
	db, err := domain.GetDB()
	if err != nil {
		return
	}

	dbItem := &domain.Item{}
	db.Find(&dbItem, ctx.Param("id"))

	if dbItem.ID == 0 {
		return "", errors.New(fmt.Sprintf(service.NotFoundF, "Item"))
	}

	dtItem := *item.CreateItem(dbItem)

	err = ctx.BindJSON(&dtItem)
	if err != nil {
		return "", err
	}

	dbItem = dtItem.ToDBO()
	db.Save(&dbItem).Commit()
	return
}

func (s Service) Delete(id string) error {
	db, err := domain.GetDB()
	if err != nil {
		return err
	}

	i := domain.Item{}
	db.Find(&i, id)

	if i.ID == 0 {
		return errors.New(fmt.Sprintf(service.NotFoundF, "Item"))
	}
	db.Unscoped().Delete(&i).Commit()

	return nil
}
