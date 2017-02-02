require "./vulnsearch/*"
require "./helper/*"
require "./model/*"
require "db"
require "kemal"

module Vulnsearch
  get "/" do
    "Hello World!"
  end

  post "/fetch" do
    dd = DownloadHelper.new
    final_year = Time.new.year

    (2002..final_year - 1).each do |year|
      dd.download(year)
    end

    dd.download(final_year)
  end
end

Kemal.run
