struct Routes::Search::View
  include Onyx::HTTP::View

  def initialize(@results : Array(Cve)); end

  html_template("search.html.ecr")
end
