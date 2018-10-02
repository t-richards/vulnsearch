require "logger"
require "option_parser"

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

  parser.on("-m up|down", "--migrate up|down", "Migrate the database") do |direction|
    Micrate.logger = logger
    Micrate::DB.connection_url = db_url

    case direction
    when "up"
      Micrate::Cli.run_up
      exit
    when "down"
      Micrate::Cli.run_down
      exit
    else
      STDERR.puts %q(Invalid direction specified. Please specify either "up" or "down")
      exit 1
    end
  rescue
    exit
  end

  parser.on("-o", "--optimize", "Optimize the database") do
    db.query("PRAGMA optimize")
    exit
  end

  parser.on("-s QUERY", "--search QUERY", "Search for things") do |query|
    pp Cve.search(query)
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
