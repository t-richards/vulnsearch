require "db"

class Cve
  DB.mapping({
    id:      String,
    summary: String,
  })
end
