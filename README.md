# vulnsearch

A fast, well-behaved replacement for other CVE search tools.

[![CircleCI](https://circleci.com/gh/t-richards/vulnsearch.svg?style=svg)](https://circleci.com/gh/t-richards/vulnsearch)

## System requirements

- [Go][golang] `~> 1.15.5`
- [SQLite][sqlite] `~> 3.33`
- Roughly 200MB of free disk space

## Getting started

```bash
# Download and install application
# Note: retry this command if you see "404 Not Found" errors
go get github.com/t-richards/vulnsearch

# Show help documentation
vulnsearch help

# Create sqlite database and migrate the schema
vulnsearch migrate

# Download CVE and product data from NVD; ~5 seconds if NVD is having a good day.
vulnsearch fetch

# Load data from compressed archives into database; ~5 minutes on a fast machine.
vulnsearch load

# Optimize the database after import; ~5 seconds
vulnsearch optimize

# Run web interface, visit http://localhost:5000/
vulnsearch
```

## Overrides

This application's database and cache file location are configurable via the environment.
By default, all files are placed in [`os.UserCacheDir()`][cachedir]

```bash
# Set custom path to the sqlite database file
VULNSEARCH_DB_PATH=/tmp/foo.sqlite3

# Set custom path to the data directory
VULNSEARCH_DATA_PATH=/opt/some/dir
```

## Hacking

```bash
# Install dependencies
make deps

# Compile application
make
```

## Contributing

1. Fork it ( <https://github.com/t-richards/vulnsearch/fork> )
2. Create your feature branch ( `git checkout -b my-new-feature` )
3. Commit your changes ( `git commit -am 'Add some feature'` )
4. Push to the branch ( `git push origin my-new-feature` )
5. Create a new Pull Request

[golang]: https://golang.org
[sqlite]: https://www.sqlite.org/index.html
[cachedir]: https://golang.org/pkg/os/#UserCacheDir
