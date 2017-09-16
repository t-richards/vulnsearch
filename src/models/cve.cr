require "db"

class Cve
  MAX_RESULTS = 1000

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

  def self.search(query)
    from_rs(VULNDB.query(default_search_query, query))
  end

  def self.default_search_query
    "SELECT * FROM cves WHERE cves MATCH ? ORDER BY rank DESC LIMIT #{MAX_RESULTS}"
  end
end
