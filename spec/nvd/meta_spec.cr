require "../spec_helper"

describe Nvd::Meta, "NVD metadata" do
  it "parses metadata properly" do
    contents = File.read("spec/fixtures/2017-meta.txt")
    meta = Nvd::Meta.parse(contents, 2017)

    # Time
    location = Time::Location.load("US/Eastern")
    expected_time = Time.new(2018, 9, 15, 3, 15, 33, location: location)

    meta.should be_a(Nvd::Meta)
    meta.last_modified.should eq(expected_time)
    meta.gz_size.should eq(7357077)
    meta.sha256.should eq("6F961E24FD62E24165054CA78CBD5A8468420C893B15E118051E607BB4E3E7D9")
    meta.year.should eq(2017)
  end
end
