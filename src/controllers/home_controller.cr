require "kemal"
require "kemal-session"

module Vulnsearch
  class HomeController
    get "/" do
      render "src/views/home/index.ecr", "src/views/layouts/default.ecr"
    end

    get "/search" do |env|
      q = "%#{env.params.query["q"]}%"
      cves = Cve.from_rs(VULNDB.query("SELECT id, summary FROM cves WHERE id LIKE ? OR summary LIKE ?", q, q))

      render "src/views/home/search.ecr", "src/views/layouts/default.ecr"
    end

    get "/fetch" do |env|
      results = env.session.string("fetch_results")
      render "src/views/home/fetch.ecr", "src/views/layouts/default.ecr"
    end

    post "/fetch" do |env|
      results = [] of String
      dd = DownloadHelper.new
      final_year = Time.new.year

      (2002..final_year - 1).each do |year|
        results << dd.download(year)
      end

      results << dd.download(final_year)
      env.session.string("fetch_results", results.join("<br />"))

      env.redirect "/fetch"
    end
  end
end
