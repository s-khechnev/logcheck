package main

import (
	"logcheck/internal/analysis/slog"
	"logcheck/internal/logcheck/englishonly"
	"logcheck/internal/logcheck/lowercase"
	nosensetivedata "logcheck/internal/logcheck/nosensitivedata"
	"logcheck/internal/logcheck/nospecnoemoji"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/multichecker"
)

var (
	slogMsgExtractor    = slog.MessagesExtractor{}
	slogVarIdsExtractor = slog.MessagesExtractor{}

	slogAnalyzers = []*analysis.Analyzer{
		lowercase.NewAnalyzer(slogMsgExtractor),
		englishonly.NewAnalyzer(slogMsgExtractor),
		nospecnoemoji.NewAnalyzer(slogMsgExtractor),
		nosensetivedata.NewAnalyzer(slogVarIdsExtractor),
	}
)

func main() {
	var analyzers []*analysis.Analyzer

	analyzers = append(analyzers, slogAnalyzers...)

	multichecker.Main(analyzers...)
}
