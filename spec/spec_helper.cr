require "spec"
require "../src/app"

# Cleans database via "DELETE" strategy
def database_cleaner
  db.exec("DELETE FROM cves")
end

Spec.before_each do
  database_cleaner
end
