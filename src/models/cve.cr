require "db"

class Cve
  DB.mapping({
    id:      String,
    summary: String,
    published: Time,
    last_modified: Time
  })

  def initialize
    @id = ""
    @summary = ""
    @published = Time.new
    @last_modified = Time.new
  end
end
