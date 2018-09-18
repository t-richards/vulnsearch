require "json"

class JsonCveItem
  JSON.mapping({
    cve:           {type: JsonCve},
    impact:        {type: JsonImpact},
    published:     {type: Time, key: "publishedDate"},
    last_modified: {type: Time, key: "lastModifiedDate"},
  })

  def id
    cve.meta.id
  rescue IndexError
    ""
  end

  def desc
    cve.description.description_data[0].value
  rescue IndexError
    ""
  end

  def cwe_id
    cve.problemtype.problemtype_data[0].description[0].value
  rescue IndexError
    ""
  end

  def cvss_v2_score : Float64
    impact.base_metric_v2.try(&.cvss_v2).try(&.base_score) || 0.0_f64
  end

  def cvss_v3_score : Float64
    impact.base_metric_v3.try(&.cvss_v3).try(&.base_severity) || 0.0_f64
  end
end

class JsonCve
  JSON.mapping({
    meta:        {type: JsonCveMeta, key: "CVE_data_meta"},
    description: {type: JsonDescription, key: "description"},
    problemtype: {type: JsonProblemType},
    affects:     {type: JsonAffects},
  })
end

class JsonProblemType
  JSON.mapping({
    problemtype_data: {type: Array(JsonProblemTypeData), key: "problemtype_data"},
  })
end

class JsonProblemTypeData
  JSON.mapping({
    description: {type: Array(JsonDescriptionData)},
  })
end

class JsonImpact
  JSON.mapping({
    base_metric_v3: {type: JsonBaseMetricV3?, key: "baseMetricV3"},
    base_metric_v2: {type: JsonBaseMetricV2?, key: "baseMetricV2"},
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
    base_severity: {type: Float64, key: "baseSeverity"},
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
    description_data: {type: Array(JsonDescriptionData), key: "description_data"},
  })
end

class JsonDescriptionData
  JSON.mapping({
    lang:  String,
    value: String,
  })
end

class JsonAffects
  JSON.mapping({
    vendor: {type: JsonVendor},
  })
end

class JsonVendor
  JSON.mapping({
    vendor_data: {type: Array(JsonVendorData)},
  })
end

class JsonVendorData
  JSON.mapping({
    vendor_name: String,
    product:     JsonProduct,
  })
end

class JsonProduct
  JSON.mapping({
    product_data: {type: Array(JsonProductData)},
  })
end

class JsonProductData
  JSON.mapping({
    product_name: String,
    version:      {type: JsonVersion},
  })
end

class JsonVersion
  JSON.mapping({
    version_data: {type: Array(JsonVersionData)},
  })
end

class JsonVersionData
  JSON.mapping({
    version_value: String,
  })
end
