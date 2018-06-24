require "db"
require "sqlite3"
require "kemal"

VULNDB = DB.open(
  ENV.fetch("DATABASE_URI", "sqlite3:./db/vulnsearch_dev.sqlite3")
)

macro render_default(filename)
  render "src/views/#{{{filename}}}.ecr", "src/views/layouts/default.ecr"
end

module Vulnsearch
  XML_DATA_DIR = "public/data"
end
