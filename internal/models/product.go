package models

import "gorm.io/gorm"

// Product is piece of software in which we are interested
type Product struct {
	gorm.Model
	Name    string
	Vendor  string
	Version string
}
