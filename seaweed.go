package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/codegangsta/cli"
	"github.com/mdb/seaweed"
	"github.com/olekukonko/tablewriter"
)

func main() {
	app := cli.NewApp()
	app.Name = "seaweed-cli"
	app.Version = "0.0.1"
	app.Usage = "Should I go surfing?"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "apiKey",
			Usage:  "Magic Seaweed API key",
			EnvVar: "MAGIC_SEAWEED_API_KEY",
		},
		cli.StringFlag{
			Name:   "cacheDir",
			Usage:  "Directory to cache API responses",
			EnvVar: "MAGIC_SEAWEED_CACHE_DIR",
		},
		cli.StringFlag{
			Name:   "cacheAge",
			Usage:  "Duration to cache API responses",
			EnvVar: "MAGIC_SEAWEED_CACHE_AGE",
		},
	}
	app.Commands = []cli.Command{
		{
			Name:        "forecast",
			Usage:       "forcast <spotId>",
			Description: "View the forecast for a spot",
			Action:      forecast,
		},
		{
			Name:        "today",
			Usage:       "today <spotId>",
			Description: "View today's forecast for a spot",
			Action:      today,
		},
		{
			Name:        "tomorrow",
			Usage:       "tomorrow <spotId>",
			Description: "View tomorrow's forecast for a spot",
			Action:      tomorrow,
		},
	}
	app.Run(os.Args)
}

func forecast(c *cli.Context) error {
	client := client(c)
	spot := c.Args().First()
	forecast, err := client.Forecast(spot)
	if err != nil {
		return err
	}

	printForecasts(spot, forecast)

	return nil
}

func today(c *cli.Context) error {
	client := client(c)
	spot := c.Args().First()
	forecast, err := client.Today(spot)
	if err != nil {
		return err
	}

	printForecasts(spot, forecast)

	return nil
}

func tomorrow(c *cli.Context) error {
	client := client(c)
	spot := c.Args().First()
	forecast, err := client.Tomorrow(spot)
	if err != nil {
		return err
	}

	printForecasts(spot, forecast)

	return nil
}

func printForecasts(spot string, forecasts []seaweed.Forecast) {
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
}

func client(c *cli.Context) *seaweed.Client {
	return &seaweed.Client{
		os.Getenv("MAGIC_SEAWEED_API_KEY"),
		&http.Client{},
		cacheAge(),
		cacheDir(),
	}
}

func cacheAge() time.Duration {
	var age time.Duration
	var _ error

	if os.Getenv("MAGIC_SEAWEED_CACHE_AGE") != "" {
		age, _ = time.ParseDuration(os.Getenv("MAGIC_SEAWEED_CACHE_DIR"))
	} else {
		age, _ = time.ParseDuration("5m")
	}

	return age
}

func cacheDir() string {
	cache := os.Getenv("MAGIC_SEAWEED_CACHE_DIR")

	if cache != "" {
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
