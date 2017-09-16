# vulnsearch

[![Build Status](https://travis-ci.org/t-richards/vulnsearch.svg?branch=master)](https://travis-ci.org/t-richards/vulnsearch)

A fast, well-behaved replacement for other CVE search tools.

## System requirements

 - [Crystal][crystal] `= 0.23.1`
 - [Shards][shards] `~> 0.7.1`
 - [libsqlite3][sqlite] `~> 3.16`
 - GNU Make

## Getting started

```bash
# Install dependencies, migrate database, compile application
# Use `make help` to get information about available make targets
$ make

# Download CVE data from NVD (`.xml.gz` files)
$ bin/vulnsearch --fetch

# Load CVE data from `.xml` files into database with fulltext indexing; ~20 seconds.
$ bin/vulnsearch --load

# Run web application
# Visit http://localhost:3000/ in your browser!
$ bin/vulnsearch
```

## Usage

TBD...

## Credits

"Protection" icon by [Chanut is Industries][chanut-is-industries], licensed under [CC BY 3.0 US][cc-by-30-us].

## Contributing

1. Fork it ( https://github.com/t-richards/vulnsearch/fork )
2. Create your feature branch ( `git checkout -b my-new-feature` )
3. Commit your changes ( `git commit -am 'Add some feature'` )
4. Push to the branch ( `git push origin my-new-feature` )
5. Create a new Pull Request

[crystal]: https://crystal-lang.org/
[shards]: https://github.com/crystal-lang/shards
[sqlite]: https://www.sqlite.org/
[chanut-is-industries]: https://thenounproject.com/chanut-is/
[cc-by-30-us]: https://creativecommons.org/licenses/by/3.0/us/
