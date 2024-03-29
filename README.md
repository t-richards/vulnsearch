# vulnsearch


A fast, offline-capable replacement for other CVE search tools.

## Screenshots

![home](https://user-images.githubusercontent.com/3905798/100484708-86068280-30cb-11eb-8f9a-a3b610e17845.png)

![search](https://user-images.githubusercontent.com/3905798/100484711-8737af80-30cb-11eb-97c0-ec5a7eba8408.png)


## System requirements

- [Go][golang] `~> 1.22`
- [SQLite][sqlite] `~> 3.33`
- Roughly 300MB of disk space

## Getting started

```bash
# Build and install application
go get -u github.com/t-richards/vulnsearch

# Show help documentation
vulnsearch help

# Create and/or update local SQLite database
vulnsearch migrate

# Download data archives from nvd.nist.gov; <5 seconds if NVD is having a good day.
vulnsearch fetch

# Load data from compressed archives into local database; <5 minutes on a fast machine with SSD.
vulnsearch load

# Optimize the local database after import; <5 seconds on a fast machine with SSD.
vulnsearch optimize

# Run web interface, visit http://localhost:5000/
vulnsearch
```

## Overrides

This application's database and cache file location are configurable via the environment.
By default, all files are placed in [`os.UserCacheDir()`][cachedir].

```bash
# Set custom path to the sqlite database file
VULNSEARCH_DB_PATH=/tmp/foo.sqlite3

# Set custom path to the data directory
VULNSEARCH_DATA_PATH=/opt/some/dir
```

## Contributing

1. Fork it ( <https://github.com/t-richards/vulnsearch/fork> )
2. Create your feature branch ( `git checkout -b my-new-feature` )
3. Commit your changes ( `git commit -am 'Add some feature'` )
4. Push to the branch ( `git push origin my-new-feature` )
5. Create a new Pull Request

## Hacking

```bash
# Build and run the application
go run main.go

# Enable sourcemaps
DEBUG=1 go run main.go
```

## License

This software is available as open source under the terms of the [MIT License][license].

[cachedir]: https://golang.org/pkg/os/#UserCacheDir
[golang]: https://golang.org
[license]: LICENSE
[sqlite]: https://www.sqlite.org/index.html
