class Cve
  DB.mapping({
    id:            String,
    description:   String,
    cwe_id:        String,
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
    @severity = ""
    @exploit_score = 0.0
    @impact_score = 0.0
    @published = Time.new
    @last_modified = Time.new
  end

  # TODO(tom): This
  def initialize(entry : XML::Node)
    @id = ""
    @description = ""
    @cwe_id = ""
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
      "INSERT INTO cves (id, description, cwe_id, severity, exploitability_score, impact_score, published, last_modified) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
      id,
      description,
      cwe_id,
      severity,
      exploit_score,
      impact_score,
      published,
      last_modified
    )
  rescue e : SQLite3::Exception
    raise e if e.code != 19
  end
end
