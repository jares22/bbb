package entity

import(
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	Name string
	Age string `valid:"int~age must be an integer"`
	CustomerID string`valid:"matches(^CM|CU\\d{8}$)~customer ID must be in the format CMxxxxxx or CUxxxxxx"`
}