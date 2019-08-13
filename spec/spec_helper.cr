require "spec"

# Use in-memory database for tests
ENV["DATABASE_URI"] ||= "sqlite3::memory:"

# Require app bits before running tests
require "../src/app"

# Cleans database via "DELETE" strategy
def database_cleaner
  db.exec("DELETE FROM cves")
  db.exec("DELETE FROM products")
  db.exec("DELETE FROM cves_products")
end

# Migrates the test database
Migrate::Migrator.new(db, logger).to_latest

# Clean database before specs
Spec.before_each do
  database_cleaner
end
