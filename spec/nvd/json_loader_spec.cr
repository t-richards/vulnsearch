require "../spec_helper"

describe Nvd::JsonLoader, "NVD JSON loading" do
  it "loads CVE entries into the database" do
    # Load fixture JSON file into database
    starting_count = Cve.count
    loader = Nvd::JsonLoader.new
    loader.load_file("spec/fixtures/nvdcve-1.0-recent.json")
    ending_count = Cve.count

    # CVE table rows should increase by 1
    (ending_count - starting_count).should eq(1)

    # Persisted CVE data should be correct
    cve = Cve.first
    cve.should_not be_nil
    cve.id.should eq("CVE-1999-1341")
    cve.description.should eq("Linux kernel before 2.3.18 or 2.2.13pre15, with SLIP and PPP options, allows local unprivileged users to forge IP packets via the TIOCSETD option on tty devices.")
    cve.cvss_v2_score.should eq(4.6)
    cve.published.should eq(Time.new(1999, 10, 22))
    cve.last_modified.should eq(Time.utc(2018, 9, 11, 14, 32, 0))
  end

  it "loads product entries into the database" do
    true.should eq(false)
  end
end
