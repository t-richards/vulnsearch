require "./vulnsearch/*"
require "./helper/*"
require "./model/*"
require "db"
require "kemal"

module Vulnsearch
  get "/" do
    render "src/views/home.ecr", "src/views/layouts/default.ecr"
  end

  post "/fetch" do
    results = [] of String
    dd = DownloadHelper.new
    final_year = Time.new.year

    (2002..final_year - 1).each do |year|
      results << dd.download(year)
    end

    results << dd.download(final_year)

    render "src/views/fetch.ecr", "src/views/layouts/default.ecr"
  end
end

Kemal.run
