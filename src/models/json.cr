class JsonCveItem
  JSON.mapping({
    cve:           {type: JsonCve},
    impact:        {type: JsonImpact},
    published:     {type: Time, key: "publishedDate"},
    last_modified: {type: Time, key: "lastModifiedDate"},
  })
end

class JsonCve
  JSON.mapping({
    meta:        {type: JsonCveMeta, key: "CVE_data_meta"},
    description: {type: JsonDescription, key: "description"},
    problemtype: {type: JsonProblemType},
  })
end

class JsonProblemType
  JSON.mapping({
    data: {type: Array(JsonProblemTypeData), key: "problemtype_data"},
  })
end

class JsonProblemTypeData
  JSON.mapping({
    description: {type: Array(JsonDescriptionData), key: "description"},
  })
end

class JsonImpact
  JSON.mapping({
    base_metric_v3:       {type: JsonBaseMetricV3?, key: "baseMetricV3"},
    base_metric_v2:       {type: JsonBaseMetricV2?, key: "baseMetricV2"},
    severity:             String?,
    exploitability_score: {type: Float64?, key: "exploitabilityScore"},
    impact_score:         {type: Float64?, key: "impactScore"},
  })
end

class JsonBaseMetricV3
  JSON.mapping({
    cvss_v3:              {type: JsonCvssV3, key: "cvssV3"},
    exploitability_score: {type: Float64, key: "exploitabilityScore"},
    impact_score:         {type: Float64, key: "impactScore"},
  })
end

class JsonCvssV3
  JSON.mapping({
    base_score:    {type: Float64, key: "baseScore"},
    base_severity: {type: String, key: "baseSeverity"},
  })
end

class JsonBaseMetricV2
  JSON.mapping({
    cvss_v2:              {type: JsonCvssV2, key: "cvssV2"},
    severity:             {type: String},
    exploitability_score: {type: Float64, key: "exploitabilityScore"},
    impact_score:         {type: Float64, key: "impactScore"},
  })
end

class JsonCvssV2
  JSON.mapping({
    base_score: {type: Float64, key: "baseScore"},
  })
end

class JsonCveMeta
  JSON.mapping({
    id: {type: String, key: "ID"},
  })
end

class JsonDescription
  JSON.mapping({
    data: {type: Array(JsonDescriptionData), key: "description_data"},
  })
end

class JsonDescriptionData
  JSON.mapping({
    lang:  String,
    value: String,
  })
end
