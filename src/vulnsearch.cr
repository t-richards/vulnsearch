require "prism"

require "./services/repo"
require "./router"
require "./actions/**"
require "./config/*"
require "./models/*"

router = Router.new
logger = Logger.new(STDOUT)
log_handler = Prism::LogHandler.new(logger)
static_handler = HTTP::StaticFileHandler.new("public", directory_listing: false)
handlers = [log_handler, static_handler, router]

host = ENV.fetch("HOST", "0.0.0.0")
port = ENV.fetch("PORT", "5000").to_i

server = Prism::Server.new(handlers, host, port, true, logger)
logger.info("Welcome to Vulnsearch v#{Vulnsearch::VERSION}")
server.listen
