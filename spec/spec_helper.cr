require "spec"

# Use in-memory database for tests
ENV["DATABASE_URI"] ||= "sqlite3:./db/vulnsearch_test.sqlite3"

# Require app bits before running tests
require "../src/app"

# Cleans database via "DELETE" strategy
def database_cleaner
  db.exec("DELETE FROM cves")
  db.exec("DELETE FROM products")
  db.exec("DELETE FROM cves_products")
end

# Migrates the test database
def migrate_test_db
  migrator = Migrate::Migrator.new(db, logger)
  migrator.to_latest
end

migrate_test_db

# Clean database before specs
Spec.before_each do
  database_cleaner
end
