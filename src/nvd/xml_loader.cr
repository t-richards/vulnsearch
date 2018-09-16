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

    def load_all_files : Int32
      db.exec("PRAGMA synchronous = OFF")
      db.exec("PRAGMA journal_mode = memory")

      @data_files.each do |file|
        logger.info "Loading data from #{file}... "
        db.exec("BEGIN TRANSACTION")
        parse_gz_data(file)
        db.exec("COMMIT")
        logger.info "Done parsing #{file}."
      end

      0
    end

    # Loads a compressed data file
    def parse_gz_data(nvdcve_file)
      File.open(nvdcve_file, "r") do |file|
        Gzip::Reader.open(file) do |inflate|
          core_parse_stuff(inflate)
        end
      end
    end

    # Loads an uncompressed data file
    def parse_data(nvdcve_file)
      File.open(nvdcve_file, "r") do |file|
        core_parse_stuff(file)
      end
    end

    def core_parse_stuff(content)
      parse_options = XML::ParserOptions.default | XML::ParserOptions::NOENT
      doc = XML.parse(content, options: parse_options)
      entries = doc.xpath_nodes(%{//*[local-name()="entry"]})
      entries.each_with_index do |entry, idx|
        # TODO(tom): Show progress bar maybe
        cve = Cve.new(entry)
        cve.save!
      end
    end
  end
end
