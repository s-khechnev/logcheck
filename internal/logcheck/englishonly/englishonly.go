package englishonly

import (
	"go/ast"
	"logcheck/internal/logcheck"
	"unicode"

	"golang.org/x/tools/go/analysis"
)

func NewAnalyzer(extractor logcheck.LogMsgExtractor) *analysis.Analyzer {
	return &analysis.Analyzer{
		Name: "englishonly",
		Doc:  "Check that the log message is english only",
		Run: func(pass *analysis.Pass) (any, error) {
			return run(extractor, pass)
		},
	}
}

func run(extractor logcheck.LogMsgExtractor, pass *analysis.Pass) (any, error) {
	for _, f := range pass.Files {
		ast.Inspect(f, func(n ast.Node) bool {
			call, ok := n.(*ast.CallExpr)
			if !ok {
				return true
			}

			msgs := extractor.ExtractLogMessages(*call, pass.TypesInfo)
			if len(msgs) == 0 {
				return true
			}

			for _, msg := range msgs {
				for _, r := range msg {
					if unicode.IsLetter(r) && !(('a' <= r && r <= 'z') || ('A' <= r && r <= 'Z')) {
						pass.Reportf(n.Pos(), "Message contains non-English letter: %s", msg)
						break
					}
				}
			}

			return true
		})
	}

	return nil, nil
}
