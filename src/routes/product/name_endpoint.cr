struct Routes::Product::NameEndpoint
  include Onyx::HTTP::Endpoint

  params do
    json require: true do
      type vendor : String
      type name : String
    end
  end

  def call
    products = ::Product.names(
      vendor: params.json.vendor.strip,
      name: params.json.name.strip
    )
    return NameView.new(products)
  end
end

Onyx::HTTP.post "/product", Routes::Product::NameEndpoint
