class Cve
  UPSERT_QUERY = <<-'EOT'
    INSERT INTO cves (id, summary, cwe_id, vendor, product, severity, exploit_score, impact_score, published, last_modified)
    VALUES           (?,  ?,       ?,      ?,      ?,       ?,        ?,             ?,            ?,         ?            )
    ON CONFLICT(id) DO UPDATE SET
      summary=excluded.summary,
      cwe_id=excluded.cwe_id,
      vendor=excluded.vendor,
      product=excluded.product,
      severity=excluded.severity,
      exploit_score=excluded.exploit_score,
      impact_score=excluded.impact_score,
      published=excluded.published,
      last_modified=excluded.last_modified
  EOT

  DB.mapping({
    id:            String,
    summary:       String,
    cwe_id:        String,
    vendor:        String,
    product:       String,
    severity:      String,
    exploit_score: Float64,
    impact_score:  Float64,
    published:     Time,
    last_modified: Time,
  })

  def initialize
    @id = ""
    @summary = ""
    @cwe_id = ""
    @vendor = ""
    @product = ""
    @severity = ""
    @exploit_score = 0.0
    @impact_score = 0.0
    @published = Time.new
    @last_modified = Time.new
  end

  # Constructs a CVE from an NVD XML "entry" node.
  def initialize(entry : XML::Node)
    namespaces = Nvd::Namespaces.namespaces

    @id = entry["id"]

    # Summary
    @summary = ""
    summary_node = entry.xpath_node("//vuln:summary", namespaces)
    if summary_node
      @summary = summary_node.inner_text
    end

    @cwe_id = ""
    @vendor = ""
    @product = ""
    @severity = ""
    @exploit_score = 0.0
    @impact_score = 0.0

    # Published
    @published = Time.new
    published_node = entry.xpath_node("//vuln:published-datetime", namespaces)
    if published_node
      @published = Time.parse_rfc3339(published_node.inner_text)
    end

    @last_modified = Time.new
  end

  # Saves the CVE to the database, or updates the data if it has changed
  def save!
    return if summary.includes?("** DISPUTED **")
    return if summary.includes?("** REJECT **")
    return if summary.includes?("** RESERVED **")

    db.exec(
      UPSERT_QUERY,
      id,
      summary,
      cwe_id,
      vendor,
      product,
      severity,
      exploit_score,
      impact_score,
      published,
      last_modified
    )
  end

  def self.first
    from_rs(db.query("SELECT * FROM cves LIMIT 1")).first
  end
end
