package slog

import (
	"logcheck/internal/logcheck/englishonly"
	"logcheck/internal/logcheck/lowercase"
	"logcheck/internal/logcheck/nosensitivedata"
	"logcheck/internal/logcheck/nospecnoemoji"

	"golang.org/x/tools/go/analysis"
)

const (
	slogName = "slog"
)

var (
	slogMsgExtractor    = MessagesExtractor{}
	slogVarIdsExtractor = VarIdsExtractor{}
)

func NewLowercaseAnalyzer() *analysis.Analyzer {
	return lowercase.NewAnalyzer(slogName, slogMsgExtractor)
}

func NewEnglishOnlyAnalyzer() *analysis.Analyzer {
	return englishonly.NewAnalyzer(slogName, slogMsgExtractor)
}

func NewNoSpecNoEmojiAnalyzer() *analysis.Analyzer {
	return nospecnoemoji.NewAnalyzer(slogName, slogMsgExtractor)
}

func NewNoSensitiveDataAnalyzer() *analysis.Analyzer {
	return nosensitivedata.NewAnalyzer(slogName, slogVarIdsExtractor)
}
