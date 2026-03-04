package main

import (
	"logcheck/internal/analysis/slog"
	"logcheck/internal/logcheck/nospecnoemoji"

	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	//singlechecker.Main(lowercase.NewAnalyzer(slog.MessagesExtractor{}))
	singlechecker.Main(nospecnoemoji.NewAnalyzer(slog.MessagesExtractor{}))
}
