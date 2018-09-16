require "xml"
require "zlib"

module Nvd
  class XmlLoader
    def initialize
      pattern = File.join(Vulnsearch::DATA_DIR, "*.xml.gz")
      @data_files = Dir.glob(pattern)
      if @data_files.size < 1
        logger.error "No data files found! Fetch data using the -f flag."
        return
      end
      @data_files.sort!.reverse!
    end

    def load_all_files
      db.exec("PRAGMA synchronous = OFF")
      db.exec("PRAGMA journal_mode = memory")
      db.exec("BEGIN TRANSACTION")
      @data_files.each do |file|
        logger.info "Loading data from #{file}... "
        parse_data(file)
        logger.info "Done parsing #{nvdcve_file}."
      end
      db.exec("COMMIT")

      0
    end

    def parse_data(nvdcve_file)
      File.open(nvdcve_file, "r") do |file|
        Gzip::Reader.open(file) do |inflate|
          parse_options = XML::ParserOptions.default | XML::ParserOptions::NOENT
          doc = XML.parse(inflate, options: parse_options)
          entries = doc.xpath_nodes("//*[local-name()=\"entry\"]")
          entries.each_with_index do |entry, idx|
            # TODO(tom): Show progress bar maybe
            cve = Cve.new(entry)
            cve.save!
          end
        end
      end
    end
  end
end
