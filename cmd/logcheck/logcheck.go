package main

import (
	"github.com/s-khechnev/logcheck/internal/config"
	"github.com/s-khechnev/logcheck/pkg/golinters/logcheck"

	"golang.org/x/tools/go/analysis/multichecker"
)

func main() {
	cfg := config.GetConfig()
	multichecker.Main(logcheck.GetAnalyzers(cfg)...)
}
