require "kemal"
require "micrate"
require "sqlite3"
require "../helpers/download_helper"
require "../helpers/file_loader"

Kemal.config.extra_options do |opts|
  opts.on("--fetch", "Fetch the latest data from NVD") do
    dd = Vulnsearch::DownloadHelper.new
    exit dd.download_all
  end

  opts.on("--load", "Load data from XML files into database") do
    loader = Vulnsearch::FileLoader.new
    exit loader.load_all_files
  end
end
