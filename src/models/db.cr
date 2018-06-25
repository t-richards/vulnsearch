require "json"

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
end
