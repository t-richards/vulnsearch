require "zlib"

module Vulnsearch
  class FileLoader
    def initialize
      pattern = File.join(Vulnsearch::DATA_DIR, "*.json.gz")
      @data_files = Dir.glob(pattern)
      if @data_files.size < 1
        STDERR.puts "No data files found! Fetch data using the -f flag."
        exit 1
      end
      @data_files.sort!.reverse
    end

    def load_all_files
      db.exec("PRAGMA synchronous = OFF")
      db.exec("PRAGMA journal_mode = memory")
      db.exec("BEGIN TRANSACTION")
      @data_files.each do |file|
        puts "Loading data from #{file}... "
        parse_data(file)
      end
      db.exec("COMMIT")

      0
    end

    def parse_data(nvdcve_file)
      File.open(nvdcve_file, "r") do |file|
        Gzip::Reader.open(file) do |inflate|
          entries = Array(JsonCveItem).from_json(inflate, root: "CVE_Items")
          entries.each_with_index do |entry, idx|
            show_progress(idx, entries.size)
            load_into_db(entry)
          end
        end
      end
      clear_line
      puts "Done."
    end

    def clear_line
      print "\r                    \r"
    end

    def show_progress(current, overall)
      clear_line
      print "#{current} of #{overall}"
    end

    def load_into_db(cve)
      desc = begin cve.cve.description.data[0].value rescue "" end
      cwe_id = begin cve.cve.problemtype.data[0].description[0].value rescue "" end

      db.exec(
        "INSERT INTO cves (id, description, cwe_id, severity, exploitability_score, impact_score, published, last_modified) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
        cve.cve.meta.id,
        desc,
        cwe_id,
        cve.impact.severity,
        cve.impact.exploitability_score,
        cve.impact.impact_score,
        cve.published,
        cve.last_modified
      )
    end
  end
end
