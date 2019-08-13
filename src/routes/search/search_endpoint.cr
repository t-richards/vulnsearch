struct Routes::Search::Endpoint
  include Onyx::HTTP::Endpoint

  params do
    query do
      type product_id : Int32
    end
  end

  def call
    results = Cve.find_by_product_id(params.query.product_id)
    return View.new(results)
  end
end

Onyx::HTTP.get "/search", Routes::Search::Endpoint
