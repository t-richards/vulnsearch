module Nvd
  class Meta
    property last_modified : Time = Time.new(1970, 1, 1)
    property gz_size : Int32 = 0
    property sha256 : String = ""

    def self.parse(data)
      new(data)
    end

    def initialize(data)
      data.split("\n").each do |line|
        next if line == ""

        parts = line.split(":", 2)
        set_part(parts[0], parts[1])
      end
    end

    def set_part(key, value)
      case key
      when "lastModifiedDate"
        @last_modified = Time.parse_rfc3339(value)
      when "gzSize"
        @gz_size = value.to_i
      when "sha256"
        @sha256 = value
      end
    end
  end
end
