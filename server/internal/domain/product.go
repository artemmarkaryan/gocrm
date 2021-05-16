package domain

type Product struct {
	ID    uint64 `gorm:"primaryKey;autoIncrement:true"`
	Name  string
	Price uint
}
