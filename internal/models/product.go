package models

import "gorm.io/gorm"

// Product is a thingy
type Product struct {
	gorm.Model
	Code  string
	Price uint
}
