require "xml"
require "zlib"

module Vulnsearch
  class FileLoader
    def initialize()
      pattern = File.join(Vulnsearch::XML_DATA_DIR, "*.xml.gz")
      @data_files = Dir.glob(pattern)
    end

    def load!()
      @data_files.each do |file|
        load_into_db(file)
      end
    end

    def load_into_db(nvdcve_file)
      parser_opts = XML::ParserOptions::RECOVER |
                    XML::ParserOptions::NONET |
                    XML::ParserOptions::NOBLANKS
      xml_doc = File.open(nvdcve_file, "r") do |file|
        Zlib::Inflate.gzip(file) do |inflate|
          XML.parse(inflate.gets_to_end, parser_opts)
        end
      end

      nvd = xml_doc.first_element_child
      if nvd
        nvd.children.each do |entry|
          puts entry["id"]
          puts entry.xpath_node("//vuln:summary")
        end
      end
    end
  end
end
