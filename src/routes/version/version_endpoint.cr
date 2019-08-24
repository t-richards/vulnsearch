struct Routes::Version::Endpoint
  include Onyx::HTTP::Endpoint

  params do
    json require: true do
      type vendor : String
      type name : String
      type version : String
    end
  end

  def call
    versions = ::Product.versions(
      vendor: params.json.vendor.strip,
      name: params.json.name.strip,
      version: params.json.version.strip
    )
    return View.new(versions)
  end
end

Onyx::HTTP.post "/version", Routes::Version::Endpoint
