require "db"

module Vulnsearch
  module Db
    class Connection
      @db : DB::Database

      def initialize(db_uri : String)
        @db = DB.open(db_uri)
      end

      def db()
        return @db
      end
    end
  end
end
