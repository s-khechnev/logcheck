package main

import (
	"logcheck/internal/analysis/slog"
	"logcheck/internal/analysis/zap"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/multichecker"
)

var slogAnalyzers = []*analysis.Analyzer{
	slog.NewLowercaseAnalyzer(),
	slog.NewEnglishOnlyAnalyzer(),
	slog.NewNoSpecNoEmojiAnalyzer(),
	slog.NewNoSensitiveDataAnalyzer(),
}

var zapAnalyzers = []*analysis.Analyzer{
	zap.NewLowercaseAnalyzer(),
	zap.NewEnglishOnlyAnalyzer(),
	zap.NewNoSpecNoEmogiAnalyzer(),
	zap.NewNoSensitiveDataAnalyzer(),
}

func main() {
	var analyzers []*analysis.Analyzer

	analyzers = append(analyzers, slogAnalyzers...)
	analyzers = append(analyzers, zapAnalyzers...)

	multichecker.Main(analyzers...)
}
