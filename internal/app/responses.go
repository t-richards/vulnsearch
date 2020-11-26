package app

// VendorResponse gives you a list of vendors
type VendorResponse struct {
	Vendors []string `json:"vendors"`
}

// ProductResponse gives you a list of product names for a given vendor
type ProductResponse struct {
	Products []string `json:"products"`
}

// VersionResponse gives you a list of versions for a given product
type VersionResponse struct {
	Versions []string `json:"versions"`
}
