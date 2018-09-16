require "../spec_helper"

Spec.before_each do
  truncate_cves
end

describe Nvd::XmlLoader, "NVD XML loading" do
  it "loads entries into the database" do
    starting_count = cve_count()

    loader = Nvd::XmlLoader.new
    loader.parse_data("spec/fixtures/nvdcve-2.0-recent.xml")

    ending_count = cve_count()

    # CVE rows should increase by 1
    (ending_count - starting_count).should eq(1)

    # Persisted CVE data should be correct
    cve = Cve.first
    cve.should_not be_nil
    cve.id.should eq("CVE-1999-1341")
  end
end
