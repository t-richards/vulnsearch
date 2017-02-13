require "../models/cve"
require "xml"
require "zlib"

module Vulnsearch
  class FileLoader
    def initialize
      pattern = File.join(Vulnsearch::XML_DATA_DIR, "*.xml.gz")
      @data_files = Dir.glob(pattern)
    end

    def load_all_files
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

      entries = xml_doc.xpath_nodes("//*[local-name()=\"entry\"]")
      entries.each do |entry|
        cve = Cve.new
        cve.id = entry["id"]
        entry.children.each do |child|
          case child.name
          when "summary"
            cve.summary = child.content
          end
        end

        puts cve.inspect
      end
    end
  end
end
