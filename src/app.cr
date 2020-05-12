# Core
require "json"
require "log"
ENV["CRYSTAL_LOG_SOURCES"] ||= "*"
Log.setup_from_env

# Shards
require "db"
require "sqlite3"

# App
require "./nvd"
require "./config/*"
require "./db/*"
require "./models/*"
require "./nvd/*"
require "./routes/*"
