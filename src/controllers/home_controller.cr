require "kemal"
require "db"
require "pg"

module Vulnsearch
  class HomeController
    get "/" do
      render "src/views/home/index.ecr", "src/views/layouts/default.ecr"
    end

    get "/search" do |env|
      q = "%#{env.params.query["q"]}%"
      cves = Cve.from_rs(VULNDB.query("SELECT id, summary FROM cves WHERE id LIKE $1", q))

      render "src/views/home/search.ecr", "src/views/layouts/default.ecr"
    end

    post "/fetch" do
      results = [] of String
      dd = DownloadHelper.new
      final_year = Time.new.year

      (2002..final_year - 1).each do |year|
        results << dd.download(year)
      end

      results << dd.download(final_year)

      render "src/views/home/fetch.ecr", "src/views/layouts/default.ecr"
    end
  end
end
