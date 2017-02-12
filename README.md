# vulnsearch

A fast, well-behaved replacement for other CVE search tools.

## System requirements

 - Crystal `~> 0.20`
 - Shards `~> 0.7`
 - libsqlite3 `~> 3.16`

## Getting started

```bash
# Install dependencies
$ shards install

# Create sqlite3 database file and load schema
$ bin/micrate up

# Run application
# Visit http://localhost:3000/ in your browser!
$ crystal run src/vulnsearch.cr
```

## Usage

TBD...

## Contributing

1. Fork it ( https://github.com/t-richards/vulnsearch/fork )
2. Create your feature branch (git checkout -b my-new-feature)
3. Commit your changes (git commit -am 'Add some feature')
4. Push to the branch (git push origin my-new-feature)
5. Create a new Pull Request
