package cmd

import (
	slog "log"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
	"github.com/mdb/seaweed-cli/internal/config"
	"github.com/mdb/seaweed-cli/internal/sw"
	"github.com/spf13/cobra"
)

var (
	Version = "dev"
)

var (
	cfgFile string

	rootCmd = &cobra.Command{
		Use:     "seaweed",
		Short:   "A Magic Seaweed CLI.",
		Version: "",
	}
)

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(
		&cfgFile,
		"config",
		"c",
		"",
		// TODO: should ~/.config/seaweed/config.yaml be the default?
		"use the specified configuration file (default is $XDG_CONFIG_HOME/seaweed/config.yml)",
	)
	rootCmd.MarkFlagFilename("config", "yaml", "yml")

	rootCmd.Version = Version
	// TODO: is this necessary?
	rootCmd.SetVersionTemplate(`seaweed {{printf "version %s\n" .Version}}`)

	rootCmd.Flags().Bool(
		"debug",
		false,
		"write debug output to debug.log",
	)

	rootCmd.Flags().BoolP(
		"help",
		"h",
		false,
		"seaweed help",
	)

	rootCmd.Run = func(cmd *cobra.Command, args []string) {
		debug, err := rootCmd.Flags().GetBool("debug")
		if err != nil {
			log.Fatal("Cannot parse debug flag", err)
		}

		p := createProgram(cfgFile, debug)

		if _, err := p.Run(); err != nil {
			log.Fatal(err)
		}
	}
}

func createProgram(configPath string, debug bool) *tea.Program {
	if debug {
		newConfigFile, err := os.OpenFile("debug.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

		if err == nil {
			log.SetOutput(newConfigFile)
			log.SetTimeFormat(time.Kitchen)
			log.SetReportCaller(true)
			log.SetLevel(log.DebugLevel)
			log.Debug("Logging to debug.log")
		} else {
			_, _ = tea.LogToFile("debug.log", "debug")
			slog.Print("Failed setting up logging", err)
		}
	}

	config, err := config.ParseConfig(cfgFile)
	if err != nil {
		log.Fatal("Error reading config", err)
	}

	return tea.NewProgram(sw.New(debug, config), tea.WithMouseCellMotion())
}
