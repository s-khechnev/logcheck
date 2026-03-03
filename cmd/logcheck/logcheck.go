package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"
	"logcheck/internal/analysis/slog"
	"logcheck/internal/logcheck/lowercase"
)

func main() {
	singlechecker.Main(lowercase.NewAnalyzer(slog.MessagesExtractor{}))
}
