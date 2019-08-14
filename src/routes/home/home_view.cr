struct Routes::Home::View
  include Onyx::HTTP::View

  def initialize(@query : String?)
  end

  def query
    HTML.escape(@query || "")
  end

  html_template("home.html.ecr")
end
