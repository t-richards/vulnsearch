require "db"

class Cve
  DB.mapping({
    id:      String,
    summary: String,
  })

  def initialize
    @id = ""
    @summary = ""
  end
end
