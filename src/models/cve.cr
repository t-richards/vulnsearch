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
end
