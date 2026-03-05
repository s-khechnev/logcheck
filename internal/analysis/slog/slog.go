package slog

import (
	"github.com/s-khechnev/logcheck/internal/logcheck/englishonly"
	"github.com/s-khechnev/logcheck/internal/logcheck/lowercase"
	"github.com/s-khechnev/logcheck/internal/logcheck/nosensitivedata"
	"github.com/s-khechnev/logcheck/internal/logcheck/nospecnoemoji"

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

func NewNoSensitiveDataAnalyzer(patterns ...string) *analysis.Analyzer {
	return nosensitivedata.NewAnalyzer(slogName, slogVarIdsExtractor, patterns...)
}
