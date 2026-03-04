package lowercase_test

import (
	"path/filepath"
	"testing"

	"github.com/s-khechnev/logcheck/internal/analysis/slog"
	"github.com/s-khechnev/logcheck/internal/analysis/zap"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestSlog(t *testing.T) {
	analysistest.Run(t, filepath.Join(analysistest.TestData(), "slog"), slog.NewLowercaseAnalyzer())
}

func TestZap(t *testing.T) {
	analysistest.Run(t, filepath.Join(analysistest.TestData(), "zap"), zap.NewLowercaseAnalyzer())
}
