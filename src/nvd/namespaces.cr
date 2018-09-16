module Nvd
  class Namespaces
    NAMESPACES = {
      "vuln"     => "http://scap.nist.gov/schema/vulnerability/0.4",
      "cpe-lang" => "http://cpe.mitre.org/language/2.0",
    }

    def self.namespaces
      NAMESPACES
    end
  end
end
