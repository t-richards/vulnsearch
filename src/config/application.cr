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

  EXAMPLES = {
    "heartbleed" => %q("CVE-2014-0160"),
    "wordpress" => %q(WordPress AND "before 4."),
    "xss" => %q("cross-site scripting"),
  }
end
