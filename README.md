# vulnsearch

[![go](https://github.com/t-richards/vulnsearch/actions/workflows/go.yml/badge.svg)](https://github.com/t-richards/vulnsearch/actions/workflows/go.yml)

A fast, offline-capable replacement for other CVE search tools.

## Screenshots

<img width="1676" height="1048" alt="Home" src="https://github.com/user-attachments/assets/d8950831-6add-42fe-abbf-920dc7ade825" />

<img width="1676" height="1048" alt="Search" src="https://github.com/user-attachments/assets/ba5ec405-1aa0-4a41-a438-6a2df0456b51" />

## System requirements

- [Go][golang] `>= 1.27`
- Roughly 1GB of disk space for all years of CVE data

## Getting started

```bash
# Build and install application
go install github.com/t-richards/vulnsearch@latest

# Show help documentation
vulnsearch help

# Download CVE feeds from nvd.nist.gov (all years, or just the given years)
vulnsearch fetch
vulnsearch fetch 2025 2026

# Load downloaded feeds into the local database (all downloaded years, or just the given years)
vulnsearch load

# Optimize the local database after loading
vulnsearch optimize

# Run web interface, visit http://localhost:3000/
vulnsearch
```

Re-running `fetch` only downloads feeds which have changed since the last
download, and re-running `load` replaces existing data for each CVE.

Databases created by earlier versions of vulnsearch are not compatible; run
`vulnsearch clobber` before loading.

## How matching works

Data comes from the [NVD CVE JSON 2.0 feeds][feeds]. Each CVE lists CPE match
criteria for the products it affects, either as an exact version
(`apache:http_server:2.4.49`), or as a wildcard version with a range
(`apache:http_server:*`, `>= 2.4.0, < 2.4.50`). Most CVEs use ranges.

Vulnsearch stores each vulnerable criterion as-is, and evaluates version ranges
when you search, so any version can be looked up, even one NVD has never listed
explicitly. Criteria which only describe the platform a vulnerable product runs
on (`"vulnerable": false`) are not treated as affected products.

Vendor version strings follow no single scheme, so version comparison is a
best-effort ordering: digit runs compare numerically (`2.4.9 < 2.4.10`),
pre-release labels sort first (`1.0rc1 < 1.0`), and other letter suffixes sort
after (`1.0.2 < 1.0.2k`). See [`internal/cpe`](internal/cpe) for details.

## Overrides

This application's database and cache file locations are configurable via the
environment. By default, all files are placed in [`os.UserCacheDir()`][cachedir].

```bash
# Set custom path to the sqlite database file
VULNSEARCH_DB_PATH=/tmp/foo.sqlite3

# Set custom path to the data directory
VULNSEARCH_DATA_PATH=/opt/some/dir

# Set the web interface port
PORT=3000
```

## Hacking

```bash
# Install code generation tools (sqlc)
script/setup

# Regenerate sqlc queries (committed)
script/generate

# Run tests, linter
script/test
script/lint

# Build and run the application
go run .
```

## License

This software is available as open source under the terms of the [MIT License][license].

[cachedir]: https://pkg.go.dev/os#UserCacheDir
[feeds]: https://nvd.nist.gov/vuln/data-feeds
[golang]: https://go.dev
[license]: LICENSE
