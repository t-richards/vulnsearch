module Vulnsearch
  VERSION = {{ `shards version`.chomp.stringify }}
  DATA_DIR = "data"
  EXAMPLES = {
    "heartbleed" => %q("CVE-2014-0160"),
    "wordpress"  => %q(WordPress AND "before 4."),
    "xss"        => %q("cross-site scripting"),
  }
  DB_URL = ENV.fetch("DATABASE_URI", "sqlite3:./db/vulnsearch_dev.sqlite3")
  VULNDB = DB.open(DB_URL)
  LOGGER = Logger.new(STDOUT)
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
