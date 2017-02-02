require "db"

class Cve
  DB.mapping({
    title: String
  })
end
