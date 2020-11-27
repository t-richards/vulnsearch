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
	p.FullName = fmt.Sprintf("%v %v %v", p.Product.Name, p.Product.Vendor, p.Product.Version)
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

var productSrc = `
{{ define "content" }}
<div class="container">
  <h1>{{ .FullName }}</h1>
  <h3>CVE List ({{ .CveCount }})</h3>

  <table class="table table-striped">
    <tr>
      <th scope="col">CVE ID</th>
      <th scope="col">Description</th>
      <th scope="col">CWE ID</th>
      <th scope="col">CVSS v2 Score</th>
      <th scope="col">CVSS v3 Score</th>
    </tr>
    {{ range .Cves }}
    <tr>
      <td>{{ .ID }}</td>
      <td>{{ .Description }}</td>
      <td>{{ .CweID }}</td>
      <td>{{ .CvssV2Score }}</td>
      <td>{{ .CvssV3Score }}</td>
    </tr>
    {{ end }}
  </table>

  <hr class="my-4">

  <h3>{{ .FullName }} CVEs by CVSSv3 severity</h3>
  {{ with .CountsV3 }}
  <table class="table table-striped">
    <tr>
      <th scope="col">Severity</th>
      <th scope="col" class="text-right">Score Range</th>
      <th scope="col" class="text-right">Count</th>
    </tr>
    <tr>
      <td>Critical</td>
      <td class="text-right">9.0 - 10.0</td>
      <td class="text-right">{{ .Critical }}</td>
    </tr>
    <tr>
      <td>High</td>
      <td class="text-right">7.0 - 8.9</td>
      <td class="text-right">{{ .High }}</td>
    </tr>
    <tr>
      <td>Medium</td>
      <td class="text-right">4.0 - 6.9</td>
      <td class="text-right">{{ .Medium }}</td>
    </tr>
    <tr>
      <td>Low</td>
      <td class="text-right">0.1 - 3.9</td>
      <td class="text-right">{{ .Low }}</td>
    </tr>
    <tr>
      <td>None</td>
      <td class="text-right">0.0</td>
      <td class="text-right">{{ .None }}</td>
    </tr>
    {{ end }}
    <tr><td colspan="3">&nbsp;</td></tr>
    <tr>
      <td class="font-weight-bold">Total</td>
      <td class="text-right"></td>
      <td class="text-right">{{ .CveCount }}</td>
    </tr>
  </table>

  <h3>{{ .FullName }} CVEs by CVSSv2 severity</h3>
  {{ with .CountsV2 }}
  <table class="table table-striped">
    <tr>
      <th scope="col">Severity</th>
      <th scope="col" class="text-right">Score Range</th>
      <th scope="col" class="text-right">Count</th>
    </tr>
    <tr>
      <td>High</td>
      <td class="text-right">7.0 - 10.0</td>
      <td class="text-right">{{ .High }}</td>
    </tr>
    <tr>
      <td>Medium</td>
      <td class="text-right">4.0 - 6.9</td>
      <td class="text-right">{{ .Medium }}</td>
    </tr>
    <tr>
      <td>Low</td>
      <td class="text-right">0.0 - 3.9</td>
      <td class="text-right">{{ .Low }}</td>
    </tr>
    {{ end }}
    <tr><td colspan="3">&nbsp;</td></tr>
    <tr>
      <td class="font-weight-bold">Total</td>
      <td class="text-right"></td>
      <td class="text-right">{{ .CveCount }}</td>
    </tr>
  </table>
</div>
{{ end }}
`
