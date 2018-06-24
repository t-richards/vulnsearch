require "core"
require "core/logger/io"
require "sqlite3"

class PG::Driver; end

uri = ENV.fetch("DATABASE_URI", "sqlite3:./db/vulnsearch_dev.sqlite3")
query_logger = Core::Logger::IO.new(STDOUT)
Services::Repository.init(uri, query_logger)
