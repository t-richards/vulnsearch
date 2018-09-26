class CveProduct < ApplicationRecord
  DB.mapping({
    cve_id:     String,
    product_id: Int64,
  })

  def initialize(@cve_id, @product_id); end

  def save!
    db.exec(
      CveProduct.insert_query,
      cve_id,
      product_id
    )
  end
end
