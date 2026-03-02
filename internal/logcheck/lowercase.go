package logcheck

import (
	"go/ast"
	"golang.org/x/tools/go/analysis"
	"unicode"
)

func NewLowercaseAnalyzer(extractor LogMsgExtractor) *analysis.Analyzer {
	return &analysis.Analyzer{
		Name: "lowercase",
		Doc:  "Check that the log message is lowercase",
		Run: func(pass *analysis.Pass) (any, error) {
			return run(extractor, pass)
		},
	}
}

func run(extractor LogMsgExtractor, pass *analysis.Pass) (any, error) {
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

			start := msgs[0]
			if len(start) > 0 && unicode.IsUpper(rune(start[0])) {
				pass.Reportf(n.Pos(), "Message starts with capital letter: %s", msgs[0])
			}

			return true
		})
	}

	return nil, nil
}
