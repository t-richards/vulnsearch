require "sqlite3"
require "core"
require "core/logger/standard"

# This service provides a global access to Core's Repository.
class Services::Repository
  class_getter! instance : Core::Repository?

  def self.init(database_url : String, logger : Core::Logger)
    @@instance = Core::Repository.new(
      DB.open(database_url),
      logger,
    )
  end
end

# A globally accessible repository instance (essentialy a shortcut).
def repo
  Services::Repository.instance
end
