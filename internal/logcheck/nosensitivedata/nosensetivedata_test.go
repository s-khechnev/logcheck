package nosensetivedata

import (
	"logcheck/internal/analysis/slog"
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestSlog(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), NewAnalyzer(slog.MessagesExtractor{}))
}
