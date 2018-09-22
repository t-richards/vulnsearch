require "zlib"

module Nvd
  class JsonLoader
    def initialize
      pattern = File.join(Vulnsearch::DATA_DIR, "*.json.gz")
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
        logger.info "Loading data from #{file}..."
        elapsed = Time.measure do
          db.exec("BEGIN TRANSACTION")
          load_gz_file(file)
          db.exec("COMMIT")
        end
        logger.info "Loaded #{file} in #{elapsed}."
      end

      0
    end

    # Loads a compressed data file
    def load_gz_file(nvdcve_file)
      File.open(nvdcve_file, "r") do |file|
        Gzip::Reader.open(file) do |inflate|
          core_parse_stuff(inflate)
        end
      end
    end

    # Loads an uncompressed data file
    def load_file(nvdcve_file)
      File.open(nvdcve_file, "r") do |file|
        core_parse_stuff(file)
      end
    end

    def core_parse_stuff(content)
      entries = Array(Vulnsearch::Json::CveItem).from_json(content, root: "CVE_Items")
      entries.each_with_index do |entry, idx|
        # TODO(tom): Show progress bar maybe
        cve = Cve.new(entry)
        cve.save!

        entry.cve.affects.vendor.vendor_data.each do |vendor_data|
          vendor_data.product.product_data.each do |product_data|
            product_data.version.version_data.each do |version_data|
              product = Product.new(
                product_data.product_name,
                vendor_data.vendor_name,
                version_data.version_value
              )
              product.save!
            end
          end
        end
      end
    end
  end
end
