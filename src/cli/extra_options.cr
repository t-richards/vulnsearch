require "kemal"
require "../helpers/download_helper"
require "../helpers/file_loader"

Kemal.config.extra_options do |opts|
  opts.on("--migrate", "Migrate the database to the latest schema") do
    VULNDB.exec("CREATE TABLE cves (id TEXT PRIMARY KEY, summary TEXT)")
    exit 0
  end

  opts.on("--fetch", "Fetch the latest data from NVD") do
    dd = Vulnsearch::DownloadHelper.new
    exit dd.download_all()
  end

  opts.on("--load", "Load data from XML files into database") do
    loader = Vulnsearch::FileLoader.new
    exit loader.load_all_files()
  end
end
