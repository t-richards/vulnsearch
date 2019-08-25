class Product < ApplicationRecord
  DB.mapping({
    id:      Int64?,
    name:    String,
    vendor:  String,
    version: String,
  })

  def initialize
    @name = ""
    @vendor = ""
    @version = ""
  end

  def initialize(@name : String, @vendor : String, @version : String)
  end

  # Saves the product to the database
  def save!
    db.exec(
      Product.insert_query,
      @id,
      @name,
      @vendor,
      @version
    )
  rescue e : SQLite3::Exception
    # Skip insert when UNIQUE constraint is violated
    error = LibSQLite3::Code.new(e.code)
    raise e unless error.constraint?
  end

  # Vendors by prefix
  def self.vendors(prefix) : Array(String)
    prefix += "%"

    rs = db.query("SELECT DISTINCT vendor FROM #{table_name} WHERE vendor LIKE ? ORDER BY vendor ASC LIMIT 100", prefix)
    flatten_resultset(rs)
  end

  # Searches for product names by vendor
  def self.names(name : String, vendor : String) : Array(String)
    name += "%"

    rs = db.query("SELECT DISTINCT name FROM #{table_name} WHERE vendor = ? AND name LIKE ? ORDER BY name ASC LIMIT 100", vendor, name)
    flatten_resultset(rs)
  end

  # Gets a list of valid product versions
  def self.versions(name : String, vendor : String, version : String)
    version += "%"

    rs = db.query("SELECT DISTINCT version FROM #{table_name} WHERE vendor = ? AND name = ? AND VERSION LIKE ? ORDER BY version ASC LIMIT 100", vendor, name, version)
    flatten_resultset(rs)
  end

  # Finds a product by key parts
  def self.find_by_parts(vendor : String, name : String, version : String)
    from_rs(db.query("SELECT * FROM #{table_name} WHERE vendor = ? AND name = ? AND version = ?", vendor, name, version)).first
  end

  # Finds a product by id
  def self.find(id : Int32)
    from_rs(db.query("SELECT * FROM #{table_name} WHERE id = ?", id)).first
  end

  CVES_RANGE_QUERY = <<-EOT
    SELECT COUNT(*) as count FROM cves
    JOIN cves_products ON cves.id = cves_products.cve_id
    WHERE cves_products.product_id = ?
    AND cves.cvss_v3_score >= ?
    AND cves.cvss_v3_score <= ?;
  EOT

  def critical_cve_count
    db.query_one(CVES_RANGE_QUERY, @id, 9.0, 10.0, as: Int32)
  end

  def high_cve_count
    db.query_one(CVES_RANGE_QUERY, @id, 7.0, 8.9, as: Int32)
  end

  def medium_cve_count
    db.query_one(CVES_RANGE_QUERY, @id, 4.0, 6.9, as: Int32)
  end

  def low_cve_count
    db.query_one(CVES_RANGE_QUERY, @id, 0.1, 3.9, as: Int32)
  end
end
