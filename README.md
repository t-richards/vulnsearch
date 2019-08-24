# vulnsearch

A fast, well-behaved replacement for other CVE search tools.

[![CircleCI](https://circleci.com/gh/t-richards/vulnsearch.svg?style=svg)](https://circleci.com/gh/t-richards/vulnsearch)

## System requirements

- [Crystal][crystal] `= 0.30.1`
- [Shards][shards] `= 0.9.0`
- [SQLite][sqlite] `~> 3.29.0`
- [TypeScript][typescript] `~> 3.5.3`

## Getting started

```bash
# Install dependencies, compile application
shards build --release

# Migrate database
bin/vulnsearch --migrate

# Download CVE data from NVD (`.json.gz` files); ~5 seconds if NVD is having a good day.
bin/vulnsearch --fetch

# Load CVE data from `.json.gz` files into database; ~3 minutes on a fast machine.
bin/vulnsearch --load
```

## Web app usage

```bash
# Build assets
cd public && tsc

# Run server, visit http://localhost:5000/
bin/vulnsearch --serve
```

## Contributing

1. Fork it ( <https://github.com/t-richards/vulnsearch/fork> )
2. Create your feature branch ( `git checkout -b my-new-feature` )
3. Commit your changes ( `git commit -am 'Add some feature'` )
4. Push to the branch ( `git push origin my-new-feature` )
5. Create a new Pull Request

[crystal]: https://crystal-lang.org/
[shards]: https://github.com/crystal-lang/shards
[sqlite]: https://www.sqlite.org/
[typescript]: https://www.typescriptlang.org/
