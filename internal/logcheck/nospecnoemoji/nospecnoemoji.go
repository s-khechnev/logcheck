package nospecnoemoji

import (
	"go/ast"
	"logcheck/internal/logcheck"
	"unicode"

	"golang.org/x/tools/go/analysis"
)

func NewAnalyzer(extractor logcheck.LogMsgExtractor) *analysis.Analyzer {
	return &analysis.Analyzer{
		Name: "nospecemoji",
		Doc:  "Check that the log message doesn't contains emoji or special chars",
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
					if !unicode.IsLetter(r) && !unicode.IsDigit(r) && !unicode.IsSpace(r) {
						pass.Reportf(n.Pos(), "Message contains special char or emoji: %s", msg)
						break
					}
				}
			}

			return true
		})
	}

	return nil, nil
}
