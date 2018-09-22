class Product < ApplicationRecord
  DB.mapping({
    id:      Int64,
    name:    String,
    version: String,
  })

  def initialize
    @id = -1
    @name = ""
    @version = ""
  end
end
