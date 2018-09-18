require "json"

module Vulnsearch::Json
  class CveItem
    JSON.mapping({
      cve:           {type: Cve},
      impact:        {type: Impact},
      published:     {type: Time, key: "publishedDate"},
      last_modified: {type: Time, key: "lastModifiedDate"},
    })

    def id : String
      cve.meta.id
    rescue IndexError
      ""
    end

    def desc : String
      cve.description.description_data[0].value
    rescue IndexError
      ""
    end

    def cwe_id : String
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

  class Cve
    JSON.mapping({
      meta:        {type: CveMeta, key: "CVE_data_meta"},
      description: {type: Description, key: "description"},
      problemtype: {type: ProblemType},
      affects:     {type: Affects},
    })
  end

  class ProblemType
    JSON.mapping({
      problemtype_data: {type: Array(ProblemTypeData), key: "problemtype_data"},
    })
  end

  class ProblemTypeData
    JSON.mapping({
      description: {type: Array(DescriptionData)},
    })
  end

  class Impact
    JSON.mapping({
      base_metric_v3: {type: BaseMetricV3?, key: "baseMetricV3"},
      base_metric_v2: {type: BaseMetricV2?, key: "baseMetricV2"},
    })
  end

  class BaseMetricV3
    JSON.mapping({
      cvss_v3:              {type: CvssV3, key: "cvssV3"},
      exploitability_score: {type: Float64, key: "exploitabilityScore"},
      impact_score:         {type: Float64, key: "impactScore"},
    })
  end

  class CvssV3
    JSON.mapping({
      base_score:    {type: Float64, key: "baseScore"},
      base_severity: {type: Float64, key: "baseSeverity"},
    })
  end

  class BaseMetricV2
    JSON.mapping({
      cvss_v2:              {type: CvssV2, key: "cvssV2"},
      severity:             {type: String},
      exploitability_score: {type: Float64, key: "exploitabilityScore"},
      impact_score:         {type: Float64, key: "impactScore"},
    })
  end

  class CvssV2
    JSON.mapping({
      base_score: {type: Float64, key: "baseScore"},
    })
  end

  class CveMeta
    JSON.mapping({
      id: {type: String, key: "ID"},
    })
  end

  class Description
    JSON.mapping({
      description_data: {type: Array(DescriptionData), key: "description_data"},
    })
  end

  class DescriptionData
    JSON.mapping({
      lang:  String,
      value: String,
    })
  end

  class Affects
    JSON.mapping({
      vendor: {type: Vendor},
    })
  end

  class Vendor
    JSON.mapping({
      vendor_data: {type: Array(VendorData)},
    })
  end

  class VendorData
    JSON.mapping({
      vendor_name: String,
      product:     Product,
    })
  end

  class Product
    JSON.mapping({
      product_data: {type: Array(ProductData)},
    })
  end

  class ProductData
    JSON.mapping({
      product_name: String,
      version:      {type: Version},
    })
  end

  class Version
    JSON.mapping({
      version_data: {type: Array(VersionData)},
    })
  end

  class VersionData
    JSON.mapping({
      version_value: String,
    })
  end
end
