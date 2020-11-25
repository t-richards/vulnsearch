package nvd

import (
	"strings"
	"time"
)

type SpecialTime struct {
	time.Time
}

func (t *SpecialTime) UnmarshalJSON(buf []byte) error {
	rawValue := strings.Trim(string(buf), `"`)
	tt, err := time.Parse("2006-01-02T15:04Z", rawValue)
	if err != nil {
		return err
	}
	t.Time = tt
	return nil
}

type Archive struct {
	Type      string      `json:"CVE_data_type"`
	Format    string      `json:"CVE_data_format"`
	Version   string      `json:"CVE_data_version"`
	CveCount  string      `json:"CVE_data_numberOfCVEs"`
	Timestamp SpecialTime `json:"CVE_data_timestamp"`
	CveItems  []CveItem   `json:"CVE_Items"`
}

type CveItem struct {
	Cve            Cve            `json:"cve"`
	Configurations Configurations `json:"configurations"`
	Impact         Impact         `json:"impact"`
	Published      SpecialTime    `json:"publishedDate"`
	LastModified   SpecialTime    `json:"lastModifiedDate"`
}

type Configurations struct {
	CveDataVersion string  `json:"CVE_data_version"`
	Nodes          []Nodes `json:"nodes"`
}

type Cve struct {
	Type           string         `json:"data_type"`
	Format         string         `json:"data_format"`
	Version        string         `json:"data_version"`
	CveMeta        CveMeta        `json:"CVE_data_meta"`
	ProblemType    ProblemType    `json:"problemtype"`
	References     References     `json:"references"`
	CveDescription CveDescription `json:"description"`
}

type CveMeta struct {
	ID       string `json:"ID"`
	Assigner string `json:"ASSIGNER"`
}

type ProblemType struct {
	ProblemTypeData []Problem `json:"problemtype_data"`
}

type Problem struct {
	Description []Description `json:"description"`
}

type References struct {
	ReferenceData []ReferenceData `json:"reference_data"`
}

type ReferenceData struct {
	URL    string   `json:"url"`
	Name   string   `json:"name"`
	Source string   `json:"refsource"`
	Tags   []string `json:"tags"`
}

type CveDescription struct {
	DescriptionData []Description `json:"description_data"`
}

type Description struct {
	Lang  string `json:"lang"`
	Value string `json:"value"`
}

type Nodes struct {
	Operator string     `json:"operator"`
	CpeMatch []CpeMatch `json:"cpe_match"`
}

type CpeMatch struct {
	Vulnerable bool   `json:"vulnerable"`
	Cpe23Url   string `json:"cpe23uri"`
}

type Impact struct {
	BaseMetricV3 BaseMetricV3 `json:"baseMetricV3"`
	BaseMetricV2 BaseMetricV2 `json:"baseMetricV2"`
}

type BaseMetricV3 struct {
	CvssV3              CvssV3  `json:"cvssv3"`
	ExploitabilityScore float32 `json:"exploitabilityScore"`
	ImpactScore         float32 `json:"impactScore"`
}

type CvssV3 struct {
	Version               string  `json:"version"`
	VectorString          string  `json:"vectorString"`
	AttackVector          string  `json:"attackVector"`
	AttackComplexity      string  `json:"attackComplexity"`
	PrivilegesRequired    string  `json:"privilegesRequired"`
	UserInteraction       string  `json:"userInteraction"`
	Scope                 string  `json:"scope"`
	ConfidentialityImpact string  `json:"confidentialityImpact"`
	IntegrityImpact       string  `json:"integrityImpact"`
	AvailabilityImpact    string  `json:"availabilityImpact"`
	BaseScore             float32 `json:"baseScore"`
	BaseSeverity          string  `json:"baseSeverity"`
}

type BaseMetricV2 struct {
	CvssV2                  CvssV2  `json:"cvssv2"`
	Severity                string  `json:"severity"`
	ExploitabilityScore     float32 `json:"exploitabilityScore"`
	ImpactScore             float32 `json:"impactScore"`
	AcInsufInfo             bool    `json:"acInsufInfo"`
	ObtainAllPrivilege      bool    `json:"obtainAllPrivilege"`
	ObtainUserPrivilege     bool    `json:"obtainUserPrivilege"`
	ObtainOtherPrivilege    bool    `json:"obtainOtherPrivilege"`
	UserInteractionRequired bool    `json:"userInteractionRequired"`
}

type CvssV2 struct {
	Version               string  `json:"version"`
	VectorString          string  `json:"vectorString"`
	AccessVector          string  `json:"accessVector"`
	AccessComplexity      string  `json:"accessComplexity"`
	Authentication        string  `json:"authentication"`
	ConfidentialityImpact string  `json:"confidentialityImpact"`
	IntegrityImpact       string  `json:"integrityImpact"`
	AvailabilityImpact    string  `json:"availabilityImpact"`
	BaseScore             float32 `json:"baseScore"`
}
