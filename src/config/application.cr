require "sqlite3"

module Vulnsearch
  VERSION  = {{ `shards version`.chomp.stringify }}
  DB_URL   = ENV.fetch("DATABASE_URI", "sqlite3:./db/vulnsearch_dev.sqlite3")
  VULNDB   = DB.open(DB_URL)
  LOGGER   = Logger.new(STDOUT)
  DATA_DIR = "data"

  DEFAULT_HEADERS = HTTP::Headers{
    "User-Agent" => "Vulnsearch v#{Vulnsearch::VERSION} (+https://github.com/t-richards/vulnsearch)",
  }

  EXAMPLES = {
    "heartbleed" => %q("CVE-2014-0160"),
    "wordpress"  => %q(WordPress AND "before 4."),
    "xss"        => %q("cross-site scripting"),
  }
end

def db
  Vulnsearch::VULNDB
end

def db_url
  Vulnsearch::DB_URL
end

def logger
  Vulnsearch::LOGGER
end

def optimize
  db.exec("PRAGMA synchronous = OFF")
  db.exec("PRAGMA journal_mode = memory")
  db.exec("PRAGMA optimize")
  db.exec("VACUUM")
end
