package zap

import (
	"github.com/s-khechnev/logcheck/internal/logcheck/englishonly"
	"github.com/s-khechnev/logcheck/internal/logcheck/lowercase"
	"github.com/s-khechnev/logcheck/internal/logcheck/nosensitivedata"
	"github.com/s-khechnev/logcheck/internal/logcheck/nospecnoemoji"

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

func NewNoSensitiveDataAnalyzer(patterns ...string) *analysis.Analyzer {
	return nosensitivedata.NewAnalyzer(zapName, zapVarIdsExtractor, patterns...)
}
