package logcheck

import (
	"github.com/golangci/plugin-module-register/register"
	"github.com/s-khechnev/logcheck/internal/analysis/slog"
	"github.com/s-khechnev/logcheck/internal/analysis/zap"
	"golang.org/x/tools/go/analysis"
)

var Analyzers = []*analysis.Analyzer{
	slog.NewLowercaseAnalyzer(),
	slog.NewEnglishOnlyAnalyzer(),
	slog.NewNoSpecNoEmojiAnalyzer(),
	slog.NewNoSensitiveDataAnalyzer(),

	zap.NewLowercaseAnalyzer(),
	zap.NewEnglishOnlyAnalyzer(),
	zap.NewNoSpecNoEmojiAnalyzer(),
	zap.NewNoSensitiveDataAnalyzer(),
}

func init() {
	register.Plugin("logcheck", New)
}

type Plugin struct{}

func (*Plugin) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	return Analyzers, nil
}

func (*Plugin) GetLoadMode() string {
	return register.LoadModeTypesInfo
}

func New(_ any) (register.LinterPlugin, error) {
	return &Plugin{}, nil
}
