package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"
	"logcheck/internal/analysis/stdlog"
	"logcheck/internal/logcheck"
)

func main() {
	singlechecker.Main(logcheck.NewLowercaseAnalyzer(stdlog.NewMessagesExtractor()))
}
