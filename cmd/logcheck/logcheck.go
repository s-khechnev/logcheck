package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"
	"logcheck/internal/analysis/slog"
	"logcheck/internal/logcheck"
)

func main() {
	singlechecker.Main(logcheck.NewLowercaseAnalyzer(slog.MessagesExtractor{}))
}
