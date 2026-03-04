package zap

import (
	"logcheck/internal/logcheck/englishonly"
	"logcheck/internal/logcheck/lowercase"
	"logcheck/internal/logcheck/nosensitivedata"
	"logcheck/internal/logcheck/nospecnoemoji"

	"golang.org/x/tools/go/analysis"
)

const (
	zapName = "zap"
)

var (
	zapMsgExtractor    = MessagesExtractor{}
	zapVarIdsExtractor = VarIdsExtractor{}
)

func NewLowercaseAnalyzer() *analysis.Analyzer {
	return lowercase.NewAnalyzer(zapName, zapMsgExtractor)
}

func NewEnglishOnlyAnalyzer() *analysis.Analyzer {
	return englishonly.NewAnalyzer(zapName, zapMsgExtractor)
}

func NewNoSpecNoEmojiAnalyzer() *analysis.Analyzer {
	return nospecnoemoji.NewAnalyzer(zapName, zapMsgExtractor)
}

func NewNoSensitiveDataAnalyzer() *analysis.Analyzer {
	return nosensitivedata.NewAnalyzer(zapName, zapVarIdsExtractor)
}
