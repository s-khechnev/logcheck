package nosensitivedata_test

import (
	"logcheck/internal/analysis/slog"
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestSlog(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), slog.NewNoSensitiveDataAnalyzer())
}
