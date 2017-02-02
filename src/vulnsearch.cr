require "./vulnsearch/*"
require "./db/*"
require "./cli/*"
require "./helper/*"
require "./model/*"
require "db"
require "pg"
require "kemal"

module Vulnsearch
  _conn = Vulnsearch::Db::Connection.new(
    ENV.fetch("DATABASE_URI", "postgres://postgres@localhost:5432/vulnsearch_dev")
  )
  db = _conn.db

  get "/" do
    render "src/views/home.ecr", "src/views/layouts/default.ecr"
  end

  get "/search" do |env|
    q = env.params.query["q"]
    cves = Cve.from_rs(db.query("SELECT id, summary FROM cves WHERE id LIKE $1", q))

    render "src/views/search.ecr", "src/views/layouts/default.ecr"
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
