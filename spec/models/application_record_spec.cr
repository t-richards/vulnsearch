require "../spec_helper"

describe ApplicationRecord do
  it "infers the table name properly" do
    Cve.table_name.should eq("cves")
    Product.table_name.should eq("products")
  end
end
