require "xml"
require "zlib"

module Vulnsearch
  class FileLoader
    def initialize
      pattern = File.join(Vulnsearch::XML_DATA_DIR, "*.xml.gz")
      @data_files = Dir.glob(pattern)
    end

    def load_all_files
      load_into_db(@data_files[0])
      return 0
      @data_files.each do |file|
        load_into_db(file)
      end

      0
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
          puts entry.namespaces
          summary = entry.xpath_node("/vuln:summary")
          puts summary if summary.not_nil!
          return
        end
      end
    end
  end
end
