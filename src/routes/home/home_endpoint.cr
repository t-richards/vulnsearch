struct Routes::Home::Endpoint
  include Onyx::HTTP::Endpoint

  params do
    query do
      type query : String?
    end
  end

  def call
    return View.new(
      query: params.query.query
    )
  end
end

Onyx::HTTP.get "/", Routes::Home::Endpoint
