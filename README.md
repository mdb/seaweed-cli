[![Build Status](https://travis-ci.org/mdb/seaweed-cli.svg?branch=master)](https://travis-ci.org/mdb/seaweed-cli)

# seaweed-cli

A Golang-based command line application for fetching surf forecast data from the [Magic Seaweed API](http://magicseaweed.com/developer/forecast-api).

## Installation

1. Download the latest [release](https://github.com/mdb/seaweed-cli/releases) for your operating system.
2. Untar the download. For example: `tar -xvf seaweed_0.0.5_darwin_x86_64.tgz`
3. Move the `seaweed` executable to your `$PATH`: `mv seaweed /usr/bin/`

Alternatively, if you chose to install from source:

```
make install
```

## Usage

You'll need:

* a Magic Seaweed API Key - this can be [requested from magicseaweed.com](http://magicseaweed.com/developer/sign-up)
* the ID of the spot you'd like to query - this can be retrieved from a spot's forecast URL. For example, Ocean City, NJ's spot ID is `391`, as per its forecast URL: `http://magicseaweed.com/Ocean-City-NJ-Surf-Report/391/`

View all commands and options:

```
$ seaweed
NAME:
   seaweed-cli - Should I go surfing?

USAGE:
   seaweed [global options] command [command options] [arguments...]

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

Example usage:

```
$ seaweed today 392
+--------------------+--------------+--------------+----------------------+------------+
|        DATE        | SOLID RATING | FADED RATING | PRIMARY SWELL HEIGHT |    WIND    |
+--------------------+--------------+--------------+----------------------+------------+
| Sat 05/21 02:00 am |            0 |            1 | 2.50ft               | 12 mph E   |
+--------------------+--------------+--------------+----------------------+------------+
| Sat 05/21 05:00 am |            0 |            1 | 2.50ft               | 10 mph E   |
+--------------------+--------------+--------------+----------------------+------------+
| Sat 05/21 08:00 am |            0 |            1 | 2.00ft               | 20 mph ENE |
+--------------------+--------------+--------------+----------------------+------------+
| Sat 05/21 11:00 am |            0 |            1 | 5.00ft               | 22 mph ENE |
+--------------------+--------------+--------------+----------------------+------------+
| Sat 05/21 02:00 pm |            0 |            2 | 6.50ft               | 22 mph NE  |
+--------------------+--------------+--------------+----------------------+------------+
| Sat 05/21 05:00 pm |            1 |            1 | 7.00ft               | 19 mph NNE |
+--------------------+--------------+--------------+----------------------+------------+
| Sat 05/21 08:00 pm |            1 |            1 | 6.50ft               | 15 mph NNE |
+--------------------+--------------+--------------+----------------------+------------+
| Sat 05/21 11:00 pm |            0 |            2 | 5.50ft               | 17 mph NE  |
+--------------------+--------------+--------------+----------------------+------------+
```

## Development

Running lint, tests, etc.:

```
make
```
