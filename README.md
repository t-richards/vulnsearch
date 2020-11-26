# vulnsearch

A fast, well-behaved replacement for other CVE search tools.

[![CircleCI](https://circleci.com/gh/t-richards/vulnsearch.svg?style=svg)](https://circleci.com/gh/t-richards/vulnsearch)

## System requirements

- [Go][golang] `~> 1.15.5`
- [SQLite][sqlite] `~> 3.33`
- Make

## Getting started

```bash
# Download and install application
go get github.com/t-richards/vulnsearch

# Migrate database
vulnsearch migrate

# Download CVE data from NVD (`.json.gz` files); ~5 seconds if NVD is having a good day.
vulnsearch fetch

# Load CVE data from `.json.gz` files into database; ~5 minutes on a fast machine.
vulnsearch load

# Optimize the database after import (optional; but consider running this once)
vulnsearch optimize
```

## Run the server

```bash
# Run server, visit http://localhost:5000/
vulnsearch

# Show valid subcommands
vulnsearch --help
```

## Hacking

```bash
# Install dependencies
make deps

# Compile application
make

# Build assets
make assets
```

## Contributing

1. Fork it ( <https://github.com/t-richards/vulnsearch/fork> )
2. Create your feature branch ( `git checkout -b my-new-feature` )
3. Commit your changes ( `git commit -am 'Add some feature'` )
4. Push to the branch ( `git push origin my-new-feature` )
5. Create a new Pull Request

[golang]: https://golang.org
[sqlite]: https://www.sqlite.org/index.html
