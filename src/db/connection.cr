require "db"
require "pg"

VULNDB = DB.open(
  ENV.fetch("DATABASE_URI", "postgres://postgres@localhost:5432/vulnsearch_dev")
)
