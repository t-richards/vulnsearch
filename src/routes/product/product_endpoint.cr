struct Routes::Product::ProductEndpoint
  include Onyx::HTTP::Endpoint

  params do
    path do
      type id : Int32
    end
  end

  def call
    return "derp"
    # return ProductView.new(
    #   product: ::Product.find(params.path.id),
    #   cves: Cve.find_by_product_id(params.path.id)
    # )
  end
end

Onyx::HTTP.get "/product/:id", Routes::Product::ProductEndpoint
