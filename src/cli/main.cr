require "logger"
require "option_parser"

migrator = Migrate::Migrator.new(db, logger)

opts = OptionParser.parse! do |parser|
  parser.banner = "Usage: #{PROGRAM_NAME} <flags>"

  parser.on("-f", "--fetch", "Fetch the latest data from NVD") do
    dd = Nvd::Downloader.new
    exit dd.download_all
  end

  parser.on("-h", "--help", "Show this help") do
    puts parser
    exit
  end

  parser.on("-l", "--load", "Load data from JSON files into database") do
    loader = Nvd::JsonLoader.new
    exit loader.load_all_files
  end

  parser.on("-m", "--migrate", "Migrate the database to the latest version") do
    migrator.to_latest
    exit
  end

  parser.on("-o", "--optimize", "Optimize the database") do
    optimize
    exit
  end

  parser.on("-s QUERY", "--search QUERY", "Search for things") do |query|
    Cve.search(query).each do |cve|
      cve.to_json(STDOUT)
    end
    puts
    exit
  end

  parser.on("-v", "--version", "Show version information") do
    puts "Vulnsearch v#{Vulnsearch::VERSION}"
    exit
  end

  parser.invalid_option do |opt|
    STDERR.puts "Invalid option: #{opt}"
    STDERR.puts "Use -h for help"
    exit 1
  end
end

# No options specified
STDERR.puts opts
exit 1
