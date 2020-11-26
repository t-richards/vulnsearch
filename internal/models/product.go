package models

// Product is piece of software in which we are interested
type Product struct {
	ID      int32
	Name    string
	Vendor  string
	Version string

	Cves []Cve `gorm:"many2many:cves_products"`
}
