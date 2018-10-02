require "http/client"

module Nvd
  class Downloader
    def initialize
      @http_client = HTTP::Client.new(Nvd::BASE_URI)
      Dir.mkdir_p(Vulnsearch::DATA_DIR)
    end

    def get(request_path : String)
      @http_client.get(request_path, Vulnsearch::DEFAULT_HEADERS)
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
        last_modified = response.headers["Last-Modified"]
        if last_modified
          lmt = Time::Format::RFC_2822.parse(last_modified)
          File.utime(atime: lmt, mtime: lmt, filename: out_path)
        end
        end_time = Time.monotonic
        elapsed = end_time - start_time
        logger.info "Successfully downloaded #{out_path} in #{elapsed}"
        return true
      end

      logger.error "Download failed for #{year}: #{response.status_code}"
      return false
    end

    def needs_download?(path, year)
      # If the file doesn't exist, it needs to be downloaded
      return true unless File.readable?(path)

      # Check the metadata first
      meta = MetaDownloader.download(year)
      return false if meta.nil?

      # Download the file again if there is a newer one
      filetime = File.info(path).modification_time
      metatime = meta.last_modified
      if metatime > filetime
        logger.info "#{year} is out-of-date: #{metatime} > #{filetime}"
        return true
      end

      false
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
