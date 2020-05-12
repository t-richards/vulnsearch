module Nvd
  class Meta
    property year : Int32 = 0
    property last_modified : Time = Time::UNIX_EPOCH
    property gz_size : Int32 = 0
    property sha256 : String = ""

    def self.parse(data, year)
      new(data, year)
    end

    def initialize(data, @year)
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
        @sha256 = value.strip
      end
    end
  end
end
