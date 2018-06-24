struct Actions::Search::Index < Prism::Action
  include Params

  params do
    param :q, String
  end

  def call
    query = params[:q]
    cves = Array(Cve).new

    return render_default "search/index" if query == ""

    begin
      cves = repo.query(Cve.where("cves MATCH ? ORDER BY rank DESC LIMIT 1000"))
    rescue ex : SQLite3::Exception
      return "Something went wrong while processing your query. Please ensure that your search conforms to the <a target='_blank' href='https://sqlite.org/fts5.html#full_text_query_syntax'>SQLite FTS5 query syntax</a> and try again."
    end

    render_default "search/index"
  end
end
