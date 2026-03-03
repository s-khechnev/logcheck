package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"
	"logcheck/internal/analysis/slog"
	"logcheck/internal/logcheck/nospecemoji"
)

func main() {
	//singlechecker.Main(lowercase.NewAnalyzer(slog.MessagesExtractor{}))
	singlechecker.Main(nospecemoji.NewAnalyzer(slog.MessagesExtractor{}))
}
