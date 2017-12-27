package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/mdb/seaweed"
	"github.com/olekukonko/tablewriter"
	logging "github.com/op/go-logging"
	"github.com/urfave/cli"
)

func cacheAge() time.Duration {
	var age time.Duration
	var _ error

	if len(os.Getenv("MAGIC_SEAWEED_CACHE_AGE")) != 0 {
		age, _ = time.ParseDuration(os.Getenv("MAGIC_SEAWEED_CACHE_AGE"))
	} else {
		age, _ = time.ParseDuration("5m")
	}

	return age
}

func cacheDir() string {
	cache := os.Getenv("MAGIC_SEAWEED_CACHE_DIR")

	if len(cache) != 0 {
		return cache
	}

	return os.TempDir()
}

func concat(arr []string) string {
	var buff bytes.Buffer

	for _, elem := range arr {
		buff.WriteString(elem)
	}

	return buff.String()
}

func printTableWithHeaders(headers []string, data [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(headers)
	table.AppendBulk(data)
	table.SetRowLine(true)
	table.Render()
}

func printForecasts(spot string, forecasts []seaweed.Forecast) error {
	s := [][]string{}
	for _, each := range forecasts {
		s = append(s, []string{
			time.Unix(each.LocalTimestamp, 0).Format("Mon 01/02 03:04 pm"),
			strconv.Itoa(each.SolidRating),
			strconv.Itoa(each.FadedRating),
			concat([]string{strconv.FormatFloat(each.Swell.Components.Primary.Height, 'f', 2, 64), each.Swell.Unit}),
			concat([]string{strconv.Itoa(each.Wind.Speed), " ", each.Wind.Unit, " ", each.Wind.CompassDirection}),
		})
	}

	if len(s) != 0 {
		printTableWithHeaders([]string{
			"date",
			"solid rating",
			"faded rating",
			"primary swell height",
			"wind",
		}, s)
	} else {
		fmt.Printf("No forecast found for spot: %s\n", spot)
	}

	return nil
}

func client(c *cli.Context) *seaweed.Client {
	logLevel := logging.INFO
	debug := os.Getenv("MAGIC_SEAWEED_DEBUG")

	if len(debug) != 0 {
		logLevel = logging.DEBUG
	}

	return &seaweed.Client{
		APIKey:     os.Getenv("MAGIC_SEAWEED_API_KEY"),
		HTTPClient: &http.Client{},
		CacheAge:   cacheAge(),
		CacheDir:   cacheDir(),
		Log:        seaweed.NewLogger(logLevel),
	}
}

func validateClient(c *seaweed.Client) {
	if len(c.APIKey) == 0 {
		fmt.Println("\nPlease set the MAGIC_SEAWEED_API_KEY environment variable")
		os.Exit(1)
	}
}
