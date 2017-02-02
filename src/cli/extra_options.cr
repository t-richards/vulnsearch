require "kemal"

Kemal.config.extra_options do |opts|
  opts.on("--migrate", "Migrate the database to the latest schema") do
    VULNDB.exec("CREATE TABLE cves (id TEXT PRIMARY KEY, summary TEXT)")
    exit 0
  end

  opts.on("--version", "Prints the version of this application") do
    puts "Vulnsearch #{Vulnsearch::VERSION}"
    exit 0
  end
end
