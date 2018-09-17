# vulnsearch

A fast, well-behaved replacement for other CVE search tools.

[![CircleCI](https://circleci.com/gh/t-richards/vulnsearch.svg?style=svg)](https://circleci.com/gh/t-richards/vulnsearch)

## System requirements

 - [Crystal][crystal] `= 0.26.0`
 - [Shards][shards] `~> 0.8.1`
 - [SQLite][sqlite] `~> 3.24.0`

## Getting started

```bash
# Install dependencies, compile application
$ shards build --release

# Migrate database
$ bin/vulnsearch --migrate up

# Download CVE data from NVD (`.json.gz` files)
$ bin/vulnsearch --fetch

# Load CVE data from `.json.gz` files into database; ~2 minutes.
$ bin/vulnsearch --load
```

## Usage

```bash
# Run search queries on database contents
$ bin/vulnsearch -s <query>
```

## Contributing

1. Fork it ( https://github.com/t-richards/vulnsearch/fork )
2. Create your feature branch ( `git checkout -b my-new-feature` )
3. Commit your changes ( `git commit -am 'Add some feature'` )
4. Push to the branch ( `git push origin my-new-feature` )
5. Create a new Pull Request

[crystal]: https://crystal-lang.org/
[shards]: https://github.com/crystal-lang/shards
[sqlite]: https://www.sqlite.org/
