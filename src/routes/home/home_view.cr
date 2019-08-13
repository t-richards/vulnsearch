struct Routes::Home::View
  include Onyx::HTTP::View

  def initialize(@query : String?)
  end

  def query
    HTML.escape(@query || "")
  end

  template("home.html.ecr")
end
