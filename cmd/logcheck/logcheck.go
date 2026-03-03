package main

import (
	"logcheck/internal/analysis/slog"
	"logcheck/internal/logcheck/lowercase"

	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	//singlechecker.Main(lowercase.NewAnalyzer(slog.MessagesExtractor{}))
	singlechecker.Main(lowercase.NewAnalyzer(slog.MessagesExtractor{}))
}
