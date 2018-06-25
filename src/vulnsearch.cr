# Stdlib
require "option_parser"

# App
require "./services/repo"
require "./config/*"
require "./models/*"
require "./helpers/*"

# Parse options here
OptionParser.parse! do |parser|
  parser.on("-f", "--fetch", "Fetch the latest data from NVD") do
    dd = Vulnsearch::DownloadHelper.new
    exit dd.download_all
  end

  parser.on("-l", "--load", "Load data from XML files into database") do
    loader = Vulnsearch::FileLoader.new
    exit loader.load_all_files
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
