package models

import (
	"time"
)

// Cve is a CVE number and associated data
type Cve struct {
	ID           string
	Description  string
	CweID        string
	CvssV2Score  float64
	CvssV3Score  float64
	Published    time.Time
	LastModified time.Time
}
