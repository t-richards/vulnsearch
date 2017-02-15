module Vulnsearch
  class HomeController
    def self.index(env)
      query = env.params.query.fetch("q", "")
      render_default "home/index"
    end

    def self.search(env)
      query = env.params.query.fetch("q", "")
      like_query = "%#{query}%"
      cves = Cve.from_rs(
        VULNDB.query(Cve.default_search_query, like_query, like_query)
      )

      render_default "home/search"
    end
  end
end
