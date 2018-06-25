require "db"
require "sqlite3"

# This service provides a global access to the database
class Services::Repository
  class_getter! db : ::DB::Database

  def self.init
    uri = ENV.fetch("DATABASE_URI", "sqlite3:./db/vulnsearch_dev.sqlite3")
    @@db = DB.open(uri)
    pp @db
  end
end

def db
  Services::Repository.db
end
