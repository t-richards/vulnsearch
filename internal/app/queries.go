package app

// SearchParams is the main search payload
type SearchParams struct {
	Vendor  *string `json:"vendor"`
	Name    *string `json:"name"`
	Version *string `json:"version"`
}
