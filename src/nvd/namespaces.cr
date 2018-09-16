module Nvd
  class Namespaces
    def self.namespaces
      {
        "vuln"     => "http://scap.nist.gov/schema/vulnerability/0.4",
        "cpe-lang" => "http://cpe.mitre.org/language/2.0",
      }
    end
  end
end
