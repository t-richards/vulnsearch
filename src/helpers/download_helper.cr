require "http/client"

module Vulnsearch
  class DownloadHelper
    BASE_URI = URI.parse("https://nvd.nist.gov")

    def initialize
      @http_client = HTTP::Client.new(BASE_URI)
      Dir.mkdir_p(Vulnsearch::XML_DATA_DIR)
    end

    def download(year)
      out_path = File.join(Vulnsearch::XML_DATA_DIR, "nvdcve-2.0-#{year}.xml.gz")
      unless needs_download?(out_path)
        puts "Already downloaded #{out_path}"
        return true
      end

      request_path = "/feeds/xml/cve/nvdcve-2.0-#{year}.xml.gz"
      response = @http_client.get(request_path)
      if response.success?
        File.write(out_path, response.body)
        puts "Successfully downloaded #{out_path}"
        return false
      end

      puts "ERROR: Download failed for #{year}"
      return false
    end

    def needs_download?(path)
      !File.readable?(path)
    end

    def download_all
      final_year = Time.new.year

      (2002..final_year - 1).each do |year|
        return 1 unless download(year)
      end

      return 1 unless download(final_year)
      0
    end
  end
end
