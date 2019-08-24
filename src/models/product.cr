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
    prefix = prefix + "%"

    rs = db.query("SELECT DISTINCT vendor FROM #{table_name} WHERE vendor LIKE ? ORDER BY vendor ASC LIMIT 100", prefix)
    flatten_resultset(rs)
  end

  # Searches for product names resultsby vendor and product name prefix
  def self.search(name : String, vendor : String): Array(String)
    name = "%" + name + "%"

    rs = db.query("SELECT DISTINCT name FROM #{table_name} WHERE vendor = ? AND name LIKE ? ORDER BY name ASC LIMIT 100", vendor, name)
    flatten_resultset(rs)
  end

  # Finds a product
  def self.find(name : String, vendor : String, version : String)
    from_rs(db.query("SELECT * FROM #{table_name} WHERE name = ? AND vendor = ? AND version = ?", name, vendor, version)).first
  end
end
