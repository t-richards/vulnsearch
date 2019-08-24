struct Routes::Vendor::Endpoint
  include Onyx::HTTP::Endpoint

  params do
    json require: true do
      type vendor : String
    end
  end

  def call
    query = params.json.vendor.strip
    return View.new([] of String) if query.empty?

    vendors = ::Product.vendors(query)
    return View.new(vendors)
  end
end

Onyx::HTTP.post "/vendor", Routes::Vendor::Endpoint
