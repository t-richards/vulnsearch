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
    cve.summary.should eq("Linux kernel before 2.3.18 or 2.2.13pre15, with SLIP and PPP options, allows local unprivileged users to forge IP packets via the TIOCSETD option on tty devices.")
    cve.published.should eq(Time.new(1999, 10, 22))
  end
end
