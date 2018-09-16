class Cve
  DB.mapping({
    id:            String,
    description:   String,
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
    @description = ""
    @cwe_id = ""
    @vendor = ""
    @product = ""
    @severity = ""
    @exploit_score = 0.0
    @impact_score = 0.0
    @published = Time.new
    @last_modified = Time.new
  end

  # TODO(tom): This
  def initialize(entry : XML::Node)
    namespaces = Nvd::Namespaces.namespaces
    puts entry.xpath_nodes("//vuln:vulnerable-software-list/vuln:product", namespaces)
    @id = entry["id"]
    @description = ""
    @cwe_id = ""
    @vendor = ""
    @product = ""
    @severity = ""
    @exploit_score = 0.0
    @impact_score = 0.0
    @published = Time.new
    @last_modified = Time.new
  end

  def save!
    return if description.includes?("** DISPUTED **")
    return if description.includes?("** REJECT **")
    return if description.includes?("** RESERVED **")

    db.exec(
      "INSERT INTO cves (id, description, cwe_id, vendor, product, severity, exploit_score, impact_score, published, last_modified) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
      id,
      description,
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
