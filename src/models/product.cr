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

    results = [] of String
    db.query("SELECT DISTINCT vendor FROM #{table_name} WHERE vendor LIKE ? ORDER BY vendor ASC LIMIT 100", prefix) do |rs|
      rs.each do
        results << rs.read(String)
      end
    end

    results
  end

  # Finds a product
  def self.find(name : String, vendor : String, version : String)
    from_rs(db.query("SELECT * FROM #{table_name} WHERE name = ? AND vendor = ? AND version = ?", name, vendor, version)).first
  end
end
