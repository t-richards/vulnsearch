require "./spec_helper"

describe Vulnsearch do
  it "renders the homepage" do
    get "/"
    response.body.should contain "Vulnsearch"
    response.body.should contain "Search for CVEs"
    response.body.should contain "Tom Richards"
  end
end
