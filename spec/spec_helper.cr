require "spec"
require "../src/app"

# Cleans database via "DELETE" strategy
def database_cleaner
  db.exec("DELETE FROM cves")
end

Spec.before_each do
  database_cleaner
end

def cve_count
  db.query("SELECT COUNT(*) FROM cves") do |rs|
    rs.each do
      return rs.read(Int32)
    end
  end

  0
end
