package dto

type DTO interface {
	Serialize() (string, error)
}
