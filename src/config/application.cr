require "db"
require "sqlite3"
require "kemal"
require "kemal-session"
require "kemal-flash"

VULNDB = DB.open(
  ENV.fetch("DATABASE_URI", "sqlite3:./db/vulnsearch_dev.sqlite3")
)

Kemal::Session.config do |config|
  config.secret = ENV.fetch("SECRET_KEY_BASE", "dbfdafa6fde1eefa81b2c5f1fbc9fec7d2037902c60a00a9df3265ed3294976a73a21518c433a1bb7100e7e254c4e7090039e979d876e0f74f531f92bf22ec36")
end

macro render_default(filename)
  render "src/views/#{{{filename}}}.ecr", "src/views/layouts/default.ecr"
end

module Vulnsearch
  XML_DATA_DIR = "public/data"
end
