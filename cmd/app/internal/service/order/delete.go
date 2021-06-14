package order

import (
	"errors"
	"github.com/artemmarkaryan/gocrm/cmd/app/internal/domain"
)

func (s Service) Delete(id string) error {
	db, err := domain.GetDB()
	if err != nil {
		return err
	}

	o := domain.Order{}
	db.Find(&o, id)

	if o.ID == 0 {
		return errors.New("Orders does not exist")
	}
	db.Unscoped().Delete(&o).Commit()

	return nil
}

