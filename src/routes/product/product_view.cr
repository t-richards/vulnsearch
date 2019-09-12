struct Routes::Product::ProductView
  include Onyx::HTTP::View

  def initialize(@product : ::Product, @cves : Array(Cve))
    @full_name = "#{@product.vendor} #{@product.name} #{@product.version}"
  end

  html_template("product.html.ecr")
end
