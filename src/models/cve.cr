class Cve < ApplicationRecord
  DB.mapping({
    id:            String,
    description:   String,
    cwe_id:        String,
    cvss_v2_score: Float64,
    cvss_v3_score: Float64,
    published:     Time,
    last_modified: Time,
  })

  JSON.mapping({
    id:            String,
    description:   String,
    cwe_id:        String,
    cvss_v2_score: Float64,
    cvss_v3_score: Float64,
    published:     Time,
    last_modified: Time,
  })

  def initialize
    @id = ""
    @description = ""
    @cwe_id = ""
    @cvss_v2_score = 0.0_f64
    @cvss_v3_score = 0.0_f64
    @published = Time.new(1970, 1, 1)
    @last_modified = Time.new(1970, 1, 1)
  end

  def initialize(item : Vulnsearch::Json::CveItem)
    @id = item.id
    @description = item.desc
    @cwe_id = item.cwe_id
    @cvss_v2_score = item.cvss_v2_score
    @cvss_v3_score = item.cvss_v3_score
    @published = item.published
    @last_modified = item.last_modified
  end

  # Saves the CVE to the database, or updates the data if it has changed
  def save!
    return if description.includes?("** DISPUTED **")
    return if description.includes?("** REJECT **")
    return if description.includes?("** RESERVED **")

    db.exec(
      Cve.insert_query,
      id,
      description,
      cwe_id,
      cvss_v2_score,
      cvss_v3_score,
      published,
      last_modified
    )
  end

  def self.find_by_product_id(id : Int32)
    query = <<-EOT
      SELECT id, description, cwe_id, cvss_v2_score, cvss_v3_score, published, last_modified FROM cves
      JOIN cves_products ON cves.id = cves_products.cve_id
      WHERE cves_products.product_id = ?
      ORDER BY cves.cvss_v3_score DESC;
    EOT

    from_rs(db.query(query, id))
  end

  # Search records
  def self.search(query) : Array(Cve)
    query = "%" + query + "%"
    results = db.query("SELECT * FROM #{table_name} WHERE id LIKE ? OR description LIKE ?", query, query)
    from_rs(results)
  end
end
