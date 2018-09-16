require "spec"
require "../src/app"

def truncate_cves
  db.exec("DELETE FROM cves")
end

def cve_count
  db.query("SELECT COUNT(*) FROM cves") do |rs|
    rs.each do
      return rs.read(Int32)
    end
  end

  0
end
