require "http/client"

module Nvd
  class Downloader
    BASE_URI = URI.parse("https://nvd.nist.gov")
    HEADERS  = HTTP::Headers{
      "User-Agent" => "Vulnsearch v#{Vulnsearch::VERSION} (+https://github.com/t-richards/vulnsearch)",
    }

    def initialize
      @http_client = HTTP::Client.new(BASE_URI)
      Dir.mkdir_p(Vulnsearch::DATA_DIR)
    end

    def get(request_path : String)
      @http_client.get(request_path, HEADERS)
    end

    def download(year)
      out_path = File.join(Vulnsearch::DATA_DIR, "nvdcve-1.0-#{year}.json.gz")
      unless needs_download?(out_path, year)
        logger.info "Already downloaded #{out_path}"
        return true
      end

      logger.info "Downloading #{year} feed..."

      start_time = Time.monotonic
      response = get("/feeds/json/cve/1.0/nvdcve-1.0-#{year}.json.gz")
      if response.success?
        File.write(out_path, response.body)
        end_time = Time.monotonic
        elapsed = end_time - start_time
        logger.info "Successfully downloaded #{out_path} in #{elapsed}"
        return true
      end

      logger.error "Download failed for #{year}: #{response.status_code}"
      return false
    end

    def needs_download?(path, year)
      return true unless File.readable?(path)

      # TODO(tom): Check meta file for possible updates for `year`
    end

    def download_all
      final_year = Time.new.year

      (2002..final_year - 1).each do |year|
        download(year)
      end

      download(final_year)
      0
    end
  end
end
