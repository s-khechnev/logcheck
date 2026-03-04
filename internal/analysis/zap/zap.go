package zap

import (
	"logcheck/internal/logcheck/englishonly"
	"logcheck/internal/logcheck/lowercase"

	"golang.org/x/tools/go/analysis"
)

const (
	zapName = "zap"
)

var (
	zapMsgExtractor = MessagesExtractor{}
)

func NewLowercaseAnalyzer() *analysis.Analyzer {
	return lowercase.NewAnalyzer(zapName, zapMsgExtractor)
}

func NewEnglishOnlyAnalyzer() *analysis.Analyzer {
	return englishonly.NewAnalyzer(zapName, zapMsgExtractor)
}
