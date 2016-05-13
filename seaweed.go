package main

import (
	"os"

	"github.com/codegangsta/cli"
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
