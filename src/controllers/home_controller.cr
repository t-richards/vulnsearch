require "kemal"
require "kemal-session"

module Vulnsearch
  class HomeController
    get "/" do
      render_default "home/index"
    end

    get "/search" do |env|
      q = "%#{env.params.query["q"]}%"
      cves = Cve.from_rs(VULNDB.query("SELECT * FROM cves WHERE id LIKE ? OR summary LIKE ? ORDER BY published DESC", q, q))

      render_default "home/search"
    end
  end
end
