package dto

import (
	"errors"
	"github.com/artemmarkaryan/gocrm/cmd/app/internal/domain"
)

var ConversionError = errors.New("Can't convert")

type DBOtoDTO interface {
	Convert(dbo domain.DBO) (DTO, error)
}
