package lowercase

import (
	"fmt"
	"go/ast"
	"golang.org/x/tools/go/analysis"
	"logcheck/internal/logcheck"
	"unicode"
)

func NewAnalyzer(extractor logcheck.LogMsgExtractor) *analysis.Analyzer {
	return &analysis.Analyzer{
		Name: "lowercase",
		Doc:  "Check that the log message is lowercase",
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

			firstMsg := []rune(msgs[0])
			if len(msgs[0]) > 0 && unicode.IsUpper(firstMsg[0]) {
				fmt.Printf("%v\n", msgs)
				pass.Reportf(n.Pos(), "Message starts with capital letter: %s", msgs[0])
			}

			return true
		})
	}

	return nil, nil
}
