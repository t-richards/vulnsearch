require "logger"
require "option_parser"

OptionParser.parse! do |parser|
  parser.banner = "Vulnsearch v#{Vulnsearch::VERSION}"
  parser.on("-f", "--fetch", "Fetch the latest data from NVD") do
    dd = Vulnsearch::DownloadHelper.new
    exit dd.download_all
  end

  parser.on("-l", "--load", "Load data from XML files into database") do
    loader = Vulnsearch::FileLoader.new
    exit loader.load_all_files
  end

  parser.on("-m DIR", "--migrate DIR", "Migrate the database") do |direction|
    Micrate.logger = Logger.new(STDOUT)
    Micrate::DB.connection_url = db_url
    if direction == "up"
      Micrate::Cli.run_up
    elsif direction == "down"
      Micrate::Cli.run_down
    else
      STDERR.puts %q(Invalid direction specified. Please specify either "up" or "down")
      exit 1
    end
  end

  parser.on("-s", "--search", "Search for things") do |query|
    # Query thing here
    results = db.query("SELECT * FROM cves WHERE cves MATCH ? ORDER BY rank DESC LIMIT 1000", query)
    pp results
  end

  parser.on("-h", "--help", "Show this help") do
    puts parser
  end

  parser.invalid_option do |opt|
    STDERR.puts "Invalid option: #{opt}"
    STDERR.puts parser
    exit 1
  end
end
