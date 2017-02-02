require "kemal"

Kemal.config.extra_options do |opts|
  opts.on("--migrate", "Migrate the database to the latest schema") do
    puts "I'm migrating!"
    exit 0
  end
end
