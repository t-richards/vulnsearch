package models

// Product is piece of software impacted by one or more CVEs
type Product struct {
	Name    string
	Vendor  string
	Version string
}
