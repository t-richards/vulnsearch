struct Routes::Search::View
  include Onyx::HTTP::View

  def initialize(@results : Array(Cve)); end

  template("search.html.ecr")
end
