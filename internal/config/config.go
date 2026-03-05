package config

import (
	"flag"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

var AvailableLoggers = []string{
	"slog",
	"zap",
}

type Config struct {
	Loggers  []string `mapstructure:"loggers"`
	Patterns []string `mapstructure:"patterns"`
}

func GetConfig() *Config {
	fs := flag.NewFlagSet("config", flag.ContinueOnError)

	loggersStr := fs.String(
		"loggers",
		"",
		fmt.Sprintf("Comma separated list of loggers. Supported: %s", strings.Join(AvailableLoggers, ", ")))

	sensPatternsStr := fs.String(
		"patterns",
		"",
		fmt.Sprintf("Comma separated list of patterns for sensitive data"))

	err := fs.Parse(os.Args[1:])
	if err != nil {
		log.Fatalf("failed to parse flags: %s", err)
	}

	var myFlags []string
	fs.VisitAll(func(f *flag.Flag) {
		myFlags = append(myFlags, fmt.Sprintf("-%s", f.Name))
	})

	var newArgs []string
	for i := 0; i < len(os.Args); i++ {
		arg := os.Args[i]

		if slices.Contains(myFlags, arg) {
			i++
			continue
		}

		newArgs = append(newArgs, arg)
	}

	os.Args = newArgs

	var loggers []string
	if len(*loggersStr) > 0 && *loggersStr != "" {
		loggers = strings.Split(*loggersStr, ",")
	}

	var sensPatterns []string
	if len(*sensPatternsStr) > 0 && *sensPatternsStr != "" {
		sensPatterns = strings.Split(*sensPatternsStr, ",")
	}

	return &Config{
		Loggers:  loggers,
		Patterns: sensPatterns,
	}
}
