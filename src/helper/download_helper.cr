require "http/client"

module Vulnsearch
  class DownloadHelper
    BASE_URI = URI.parse("https://nvd.nist.gov")
    OUTPUT_DIRECTORY = "public/data"

    def initialize()
      @http_client = HTTP::Client.new(BASE_URI)
      Dir.mkdir_p(OUTPUT_DIRECTORY)
    end

    def download(year)
      out_path = File.join(OUTPUT_DIRECTORY, "nvdcve-2.0-#{year}.xml.gz")
      unless needs_download?(out_path)
        return "Already downloaded #{out_path}"
      end

      request_path = "/feeds/xml/cve/nvdcve-2.0-#{year}.xml.gz"
      response = @http_client.get(request_path)
      if response.success?
        File.write(out_path, response.body)
        return "Successfully downloaded #{out_path}"
      end

      return "ERROR: Download failed for #{year}"
    end

    def needs_download?(path)
      !File.readable?(path)
    end
  end
end
