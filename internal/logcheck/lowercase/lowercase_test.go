package lowercase

import (
	"golang.org/x/tools/go/analysis/analysistest"
	"logcheck/internal/analysis/slog"
	"testing"
)

func TestSlog(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), NewAnalyzer(slog.MessagesExtractor{}))
}
