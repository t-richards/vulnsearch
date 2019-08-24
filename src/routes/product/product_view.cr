struct Routes::Product::ProductView
  include Onyx::HTTP::View

  def initialize(@product : ::Product, @cves : Array(Cve))
  end

  html_template("product.html.ecr")
end
