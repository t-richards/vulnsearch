module Nvd
  class MetaDownloader
    def self.download(year) : Nvd::Meta?
      client = HTTP::Client.new(Nvd::BASE_URI)
      response = client.get("/feeds/json/cve/#{VERSION}/nvdcve-#{VERSION}-#{year}.meta", Vulnsearch::DEFAULT_HEADERS)
      return Nvd::Meta.parse(response.body, year) if response.success?

      nil
    end
  end
end
