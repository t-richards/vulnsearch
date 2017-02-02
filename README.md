# vulnsearch

A fast, well-behaved replacement for other CVE search tools.

## System requirements

 - Crystal `~> 0.20.5`
 - Shards `~> 0.7.1`
 - PostgreSQL `~> 9.6.1` (`postgres` user with no password)

## Getting started

```bash
# Install dependencies
$ shards install

# Compile application
$ crystal build -o bin/vulnsearch src/vulnsearch.cr

# Create required database schema
$ bin/vulnsearch --migrate

# Run application
# Visit http://localhost:3000/ in your browser!
$ bin/vulnsearch
```

## Usage

TBD...

## Contributing

1. Fork it ( https://github.com/t-richards/vulnsearch/fork )
2. Create your feature branch (git checkout -b my-new-feature)
3. Commit your changes (git commit -am 'Add some feature')
4. Push to the branch (git push origin my-new-feature)
5. Create a new Pull Request
