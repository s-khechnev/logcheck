package logcheck

import (
	"log"
	"slices"

	"github.com/golangci/plugin-module-register/register"
	"github.com/s-khechnev/logcheck/internal/analysis/slog"
	"github.com/s-khechnev/logcheck/internal/analysis/zap"
	"github.com/s-khechnev/logcheck/internal/config"
	"golang.org/x/tools/go/analysis"
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
	zap.NewNoSpecNoEmojiAnalyzer(),
	zap.NewNoSensitiveDataAnalyzer(),
}

const PluginName = "logcheck"

func init() {
	register.Plugin(PluginName, New)
}

func GetAnalyzers(cfg *config.Config) []*analysis.Analyzer {
	var analyzers []*analysis.Analyzer

	if len(cfg.Loggers) == 0 || (len(cfg.Loggers) == 1 && cfg.Loggers[0] == "") {
		analyzers = append(analyzers, slogAnalyzers...)
		analyzers = append(analyzers, zapAnalyzers...)
		return analyzers
	}

	slices.Sort(cfg.Loggers)
	slices.Compact(cfg.Loggers)

	for _, logger := range cfg.Loggers {
		if logger == "zap" {
			analyzers = append(analyzers, zapAnalyzers...)
			continue
		}

		if logger == "slog" {
			analyzers = append(analyzers, slogAnalyzers...)
			continue
		}
	}

	return analyzers
}

type Plugin struct {
	cfg *config.Config
}

func (p *Plugin) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	return GetAnalyzers(p.cfg), nil
}

func (*Plugin) GetLoadMode() string {
	return register.LoadModeTypesInfo
}

func New(conf any) (register.LinterPlugin, error) {
	cfg, err := register.DecodeSettings[config.Config](conf)
	if err != nil {
		log.Fatalf("error decoding settings: %s", err)
	}

	return &Plugin{
		cfg: &cfg,
	}, nil
}
