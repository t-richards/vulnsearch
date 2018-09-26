abstract class ApplicationRecord
  # Infers the table name with bad pluralization
  def self.table_name
    {{ @type.stringify }}.split(/(?=[A-Z][^A-Z])/).sort.map{|n| n + "s"}.join("_").downcase
  end

  # Generate insert query from properties. Used in save.
  def self.insert_query
    %<INSERT INTO #{table_name} (#{{{ @type.instance_vars.join(", ") }}}) VALUES (#{{{ @type.instance_vars.map { "?" }.join(", ") }}})>
  end

  # Retrieve first record
  def self.first
    from_rs(db.query("SELECT * FROM #{table_name} LIMIT 1")).first
  end

  # Count records
  def self.count
    db.query("SELECT COUNT(*) FROM #{table_name}") do |rs|
      rs.each do
        return rs.read(Int32)
      end
    end

    0
  end
end
