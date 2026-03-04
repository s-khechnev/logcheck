package main

import (
	"logcheck/internal/analysis/slog"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/multichecker"
)

var slogAnalyzers = []*analysis.Analyzer{
	slog.NewLowercaseAnalyzer(),
	slog.NewEnglishOnlyAnalyzer(),
	slog.NewNoSpecNoEmojiAnalyzer(),
	slog.NewNoSensitiveDataAnalyzer(),
}

func main() {
	var analyzers []*analysis.Analyzer

	analyzers = append(analyzers, slogAnalyzers...)

	multichecker.Main(analyzers...)
}
