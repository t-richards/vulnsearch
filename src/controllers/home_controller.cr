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
  end
end
