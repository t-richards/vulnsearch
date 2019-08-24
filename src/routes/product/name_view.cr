struct Routes::Product::NameView
  include Onyx::HTTP::View

  def initialize(@products : Array(String)); end

  json({
    products: @products,
  })
end
