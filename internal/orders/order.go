package orders

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	ID     uint64
	Amount float64
	Status string
}
