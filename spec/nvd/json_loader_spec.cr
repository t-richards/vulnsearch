require "../spec_helper"

describe Nvd::JsonLoader, "NVD JSON loading" do
  it "loads CVE entries into the database" do
    starting_count = cve_count()

    loader = Nvd::JsonLoader.new
    loader.load_file("spec/fixtures/nvdcve-1.0-recent.json")

    ending_count = cve_count()

    # CVE rows should increase by 1
    (ending_count - starting_count).should eq(1)

    # Persisted CVE data should be correct
    cve = Cve.first
    cve.should_not be_nil
    cve.id.should eq("CVE-1999-1341")
    cve.summary.should eq("Linux kernel before 2.3.18 or 2.2.13pre15, with SLIP and PPP options, allows local unprivileged users to forge IP packets via the TIOCSETD option on tty devices.")
    cve.severity.should eq("4.6")
    cve.published.should eq(Time.new(1999, 10, 22))
    cve.last_modified.should eq(Time.new(2018, 9, 11, 10, 32, 55, nanosecond: 857000000))
  end
end
