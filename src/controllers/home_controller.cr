require "kemal"
require "kemal-session"

module Vulnsearch
  class HomeController
    get "/" do |env|
      query = env.params.query["q"]
      render_default "home/index"
    end

    get "/search" do |env|
      query = env.params.query["q"]
      like_query = "%#{query}%"
      cves = Cve.from_rs(VULNDB.query("SELECT * FROM cves WHERE id LIKE ? OR summary LIKE ? ORDER BY published DESC", like_query, like_query))

      render_default "home/search"
    end
  end
end
