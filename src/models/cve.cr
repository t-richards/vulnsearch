class Cve
  UPSERT_QUERY = <<-'EOT'
    INSERT INTO cves (id, description, cwe_id, vendor, product, cvss_v2_score, cvss_v3_score, published, last_modified)
    VALUES           (?,  ?,           ?,      ?,      ?,       ?,             ?,             ?,         ?            )
    ON CONFLICT(id) DO UPDATE SET
      description=excluded.description,
      cwe_id=excluded.cwe_id,
      vendor=excluded.vendor,
      product=excluded.product,
      cvss_v2_score=excluded.cvss_v2_score,
      published=excluded.published,
      last_modified=excluded.last_modified
  EOT

  DB.mapping({
    id:            String,
    description:   String,
    cwe_id:        String,
    vendor:        String,
    product:       String,
    cvss_v2_score: String,
    cvss_v3_score: String,
    published:     Time,
    last_modified: Time,
  })

  def initialize
    @id = ""
    @description = ""
    @cwe_id = ""
    @vendor = ""
    @product = ""
    @cvss_v2_score = ""
    @cvss_v3_score = ""
    @published = Time.new(1970, 1, 1)
    @last_modified = Time.new(1970, 1, 1)
  end

  def initialize(item : JsonCveItem)
    @id = item.id
    @description = item.desc
    @cwe_id = item.cwe_id
    @vendor = ""
    @product = ""
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
      UPSERT_QUERY,
      id,
      description,
      cwe_id,
      vendor,
      product,
      cvss_v2_score,
      cvss_v3_score,
      published,
      last_modified
    )
  end

  def self.first
    from_rs(db.query("SELECT * FROM cves LIMIT 1")).first
  end
end
