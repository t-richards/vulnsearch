require "db"

class Cve
  DB.mapping({
    id:            String,
    summary:       String,
    cwe_id:        String,
    published:     Time,
    last_modified: Time,
  })

  def initialize
    @id = ""
    @summary = ""
    @cwe_id = ""
    @published = Time.new
    @last_modified = Time.new
  end

  def self.default_search_query
    "SELECT * FROM cves WHERE id LIKE ? OR summary LIKE ? ORDER BY id DESC LIMIT 10000"
  end
end
