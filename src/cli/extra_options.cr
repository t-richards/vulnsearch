require "kemal"
require "../helpers/download_helper"
require "../helpers/file_loader"

Kemal.config.extra_options do |opts|
  opts.on("--migrate", "Migrate the database to the latest schema") do
    VULNDB.exec("CREATE TABLE cves (id TEXT PRIMARY KEY, summary TEXT)")
    exit 0
  end

  opts.on("--fetch", "Fetch the latest data from NVD") do
    results = [] of String
    dd = Vulnsearch::DownloadHelper.new
    final_year = Time.new.year

    (2002..final_year - 1).each do |year|
      exit 1 unless dd.download(year)
    end

    exit 1 unless dd.download(final_year)
    exit 0
  end

  opts.on("--load", "Load data from XML files into database") do
    loader = Vulnsearch::FileLoader.new
    loader.load!()
    exit 0
  end
end
