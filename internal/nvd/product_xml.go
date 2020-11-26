package nvd

import (
	"encoding/xml"
)

type CpeList struct {
	XMLName xml.Name  `xml:"cpe-list"`
	Items   []CpeItem `xml:"cpe-item"`
}

type CpeItem struct {
	XMLName xml.Name `xml:"cpe-item"`
	Name    string   `xml:"name,attr"`
	Title   string   `xml:"title"`
}
