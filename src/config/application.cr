require "sqlite3"

module Vulnsearch
  VERSION      = {{ `shards version`.chomp.stringify }}
  DATABASE_URL = ENV.fetch("DATABASE_URL", "sqlite3:./db/vulnsearch_dev.sqlite3")
  VULNDB       = DB.open(DATABASE_URL)
  Log          = ::Log.for("vulnsearch")
  DATA_DIR     = "data"

  DEFAULT_HEADERS = HTTP::Headers{
    "User-Agent" => "Vulnsearch v#{Vulnsearch::VERSION} (+https://github.com/t-richards/vulnsearch)",
  }
end

def db
  Vulnsearch::VULNDB
end

def optimize
  db.exec("PRAGMA synchronous = OFF")
  db.exec("PRAGMA journal_mode = memory")
  db.exec("PRAGMA optimize")
  db.exec("VACUUM")
end
