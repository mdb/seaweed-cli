[![Build Status](https://travis-ci.org/mdb/seaweed-cli.svg?branch=master)](https://travis-ci.org/mdb/seaweed-cli)

# seaweed-cli

A Golang-based command line application for fetching surf forecast data from the [Magic Seaweed API](http://magicseaweed.com/developer/forecast-api).

## Installation

1. Download the latest [release](https://github.com/mdb/seaweed-cli/releases) for your operating system.
2. Untar the download. For example: `tar -xvf seaweed-cli_0.0.5_darwin_x86_64.tgz`
3. Move the `seaweed-cli` to your `$PATH`: `mv seaweed-cli /usr/bin/`

## Usage

You'll need:

* Magic Seaweed API Key - this can be [requested from magicseaweed.com](http://magicseaweed.com/developer/sign-up)
* the ID of the spot you'd like to query - this can be retrieved from a spot's forecast URL. For example, Ocean City, NJ's spot ID is `391`, as per its forecast URL: `http://magicseaweed.com/Ocean-City-NJ-Surf-Report/391/`

```
$ seaweed-cli
NAME:
   seaweed-cli - Should I go surfing?

USAGE:
   seaweed-cli [global options] command [command options] [arguments...]

VERSION:
   0.0.5

COMMANDS:
     forecast   forcast <spotId>
     today      today <spotId>
     tomorrow   tomorrow <spotId>
     weekend    weekend <spotId>

GLOBAL OPTIONS:
   --apiKey value       Magic Seaweed API key [$MAGIC_SEAWEED_API_KEY]
   --cacheDir value     Directory to cache API responses [$MAGIC_SEAWEED_CACHE_DIR]
   --cacheAge value     Duration to cache API responses [$MAGIC_SEAWEED_CACHE_AGE]
   --help, -h           show help
   --version, -v        print the version
```

## Development

Running lint, tests, etc.:

```
make
```
