package nosensitivedata_test

import (
	"logcheck/internal/analysis/slog"
	"path/filepath"
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestSlog(t *testing.T) {
	analysistest.Run(t, filepath.Join(analysistest.TestData(), "slog"), slog.NewNoSensitiveDataAnalyzer())
}
