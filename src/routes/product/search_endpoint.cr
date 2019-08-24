struct Routes::Product::SearchEndpoint
    include Onyx::HTTP::Endpoint

    params do
      form require: true do
        type vendor : String
        type product : String
        type version : String
      end
    end

    def call
      product_id = ::Product.find_by_parts(
        vendor: params.form.vendor.strip,
        name: params.form.product.strip,
        version: params.form.version.strip
      ).id

      redirect("/product/#{product_id}")
    end
  end

  Onyx::HTTP.post "/search", Routes::Product::SearchEndpoint
