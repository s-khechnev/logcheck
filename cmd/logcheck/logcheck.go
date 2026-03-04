package main

import (
	"logcheck/internal/analysis/slog"
	nosensetivedata "logcheck/internal/logcheck/nosensitivedata"

	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	//singlechecker.Main(lowercase.NewAnalyzer(slog.MessagesExtractor{}))
	singlechecker.Main(nosensetivedata.NewAnalyzer(slog.MessagesExtractor{}))
}
