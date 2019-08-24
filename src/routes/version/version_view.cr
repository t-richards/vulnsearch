struct Routes::Version::View
  include Onyx::HTTP::View

  def initialize(@versions : Array(String)); end

  json({
    versions: @versions,
  })
end
