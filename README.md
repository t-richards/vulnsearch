# vulnsearch

A fast, well-behaved replacement for other CVE search tools.

## System requirements

 - [Crystal][crystal] `~> 0.25.0`
 - [Shards][shards] `~> 0.8.1`
 - [SQLite][sqlite] `~> 3.24.0`

## Getting started

```bash
# Install dependencies, compile application
$ shards build

# Migrate database
$ bin/vulnsearch --migrate up

# Download CVE data from NVD (`.json.gz` files)
$ bin/vulnsearch --fetch

# Load CVE data from `.json` files into database; ~3 minutes.
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
