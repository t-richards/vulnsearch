struct Routes::Vendor::View
  include Onyx::HTTP::View

  def initialize(@vendors : Array(String)); end

  json({
    vendors: @vendors,
  })
end
