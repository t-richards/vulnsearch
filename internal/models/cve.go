package models

import "time"

// Cve is a CVE number and associated data
type Cve struct {
	ID          string
	Description string
	CweID       string
	CvssV2Score float32
	CvssV3Score float32

	Products []Product `gorm:"many2many:cves_products"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
