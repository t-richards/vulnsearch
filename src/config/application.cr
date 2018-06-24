module Vulnsearch
  VERSION = "0.1.0"
  XML_DATA_DIR = "public/data"
  EXAMPLES = {
    "heartbleed" => %q("CVE-2014-0160"),
    "wordpress"  => %q(WordPress AND "before 4."),
    "xss"        => %q("cross-site scripting"),
  }
end
