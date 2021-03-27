package views

import (
	"fmt"

	"github.com/t-richards/vulnsearch/internal/models"
)

// ProductView is the data this template expects to receive
type ProductView struct {
	Product models.Product
	Cves    []models.Cve

	// Additional Product data
	FullName string
	CveCount int

	CountsV2 CountsV2
	CountsV3 CountsV3
}

type CountsV2 struct {
	High   int
	Medium int
	Low    int
}

type CountsV3 struct {
	Critical int
	High     int
	Medium   int
	Low      int
	None     int
}

// Prepare computes some additional data for the view
func (p *ProductView) Prepare() {
	p.FullName = fmt.Sprintf("%v %v %v", p.Product.Vendor, p.Product.Name, p.Product.Version)
	p.CveCount = len(p.Cves)

	for _, cve := range p.Cves {
		// Count v3 totals
		switch true {
		case cve.CvssV3Score >= 9.0:
			p.CountsV3.Critical++
		case cve.CvssV3Score >= 7.0:
			p.CountsV3.High++
		case cve.CvssV3Score >= 4.0:
			p.CountsV3.Medium++
		case cve.CvssV3Score >= 0.1:
			p.CountsV3.Low++
		default:
			p.CountsV3.None++
		}

		// Count v2 totals
		switch true {
		case cve.CvssV2Score >= 7.0:
			p.CountsV2.High++
		case cve.CvssV2Score >= 4.0:
			p.CountsV2.Medium++
		default:
			p.CountsV2.Low++
		}
	}
}
