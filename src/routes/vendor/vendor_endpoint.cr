struct Routes::Vendor::Endpoint
  include Onyx::HTTP::Endpoint

  params do
    query do
      type prefix : String
    end
  end

  def call
    vendors = Product.vendors(params.query.prefix)
    return View.new(vendors)
  end
end

Onyx::HTTP.get "/vendor", Routes::Vendor::Endpoint
