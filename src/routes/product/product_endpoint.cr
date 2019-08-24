struct Routes::Product::Endpoint
  include Onyx::HTTP::Endpoint

  params do
    json require: true do
      type vendor : String
      type name : String
    end
  end

  def call
    products = ::Product.search(
      vendor: params.json.vendor.strip,
      name: params.json.name.strip
    )
    return View.new(products)
  end
end

Onyx::HTTP.post "/product", Routes::Product::Endpoint
