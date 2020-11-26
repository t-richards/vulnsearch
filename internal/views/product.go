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
	FullName         string
	CveCount         int
	CriticalCveCount int
	HighCveCount     int
	MediumCveCount   int
	LowCveCount      int
}

// Prepare sets some additional data on the view
func (p *ProductView) Prepare() {
	p.FullName = fmt.Sprintf("%v %v %v", p.Product.Name, p.Product.Vendor, p.Product.Version)
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
      <th scope="col">CVSS v3 Score</th>
    </tr>
    {{ range .Cves }}
    <tr>
      <td>{{ .Id }}</td>
      <td>{{ .Description }}</td>
      <td>{{ .CweID }}</td>
      <td>{{ .CvssV3Score }}</td>
    </tr>
    {{ end }}
  </table>

  <hr class="my-4">

  <h3>{{ .FullName }} CVEs by severity</h3>
  <table class="table table-striped">
    <tr>
      <th scope="col">Severity</th>
      <th scope="col" class="text-right">Score Range</th>
      <th scope="col" class="text-right">Count</th>
    </tr>
    <tr>
      <td>Critical</td>
      <td class="text-right">9.0 - 10.0</td>
      <td class="text-right">{{ .CriticalCveCount }}</td>
    </tr>
    <tr>
      <td>High</td>
      <td class="text-right">7.0 - 8.9</td>
      <td class="text-right">{{ .HighCveCount }}</td>
    </tr>
    <tr>
      <td>Medium</td>
      <td class="text-right">4.0 - 6.9</td>
      <td class="text-right">{{ .MediumCveCount }}</td>
    </tr>
    <tr>
      <td>Low</td>
      <td class="text-right">0.1 - 3.9</td>
      <td class="text-right">{{ .LowCveCount }}</td>
    </tr>
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
