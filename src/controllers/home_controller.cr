module Vulnsearch
  class HomeController
    def self.index(env)
      query = env.params.query.fetch("q", "")
      render_default "home/index"
    end

    def self.search(env)
      query = env.params.query.fetch("q", "")
      cves = Array(Cve).new

      return render_default "home/search" if query == ""

      begin
        cves = Cve.search(query)
      rescue ex : SQLite3::Exception
        return "Something went wrong while processing your query. Please ensure that your search conforms to the <a target='_blank' href='https://sqlite.org/fts5.html#full_text_query_syntax'>SQLite FTS5 query syntax</a> and try again."
      end

      render_default "home/search"
    end
  end
end
