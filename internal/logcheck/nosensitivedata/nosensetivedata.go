package nosensetivedata

import (
	"go/ast"
	"logcheck/internal/logcheck"
	"regexp"

	"golang.org/x/tools/go/analysis"
)

var sensitivePatterns = []*regexp.Regexp{
	regexp.MustCompile(`(?i)pass(word)?|pwd`),
	regexp.MustCompile(`(?i)token|jwt|refresh|access_token`),
	regexp.MustCompile(`(?i)api[_-]?key|apikey|api_secret`),
	regexp.MustCompile(`(?i)secret|private_key`),
	regexp.MustCompile(`(?i)card|ccv|cvv|credit_card`),
	regexp.MustCompile(`(?i)auth|authorization|basic_auth`),
	regexp.MustCompile(`(?i)session|cookie|session_id`),
	regexp.MustCompile(`(?i)conn(ection)?[_-]?string|dsn`),
	regexp.MustCompile(`(?i)email|e-?mail|mail`),
}

func NewAnalyzer(extractor logcheck.LogVarIdsExtractor) *analysis.Analyzer {
	return &analysis.Analyzer{
		Name: "nosensetivedata",
		Doc:  "Check that the log message doesn't contains sensitive data",
		Run: func(pass *analysis.Pass) (any, error) {
			return run(extractor, pass)
		},
	}
}

func run(extractor logcheck.LogVarIdsExtractor, pass *analysis.Pass) (any, error) {
	for _, f := range pass.Files {
		ast.Inspect(f, func(n ast.Node) bool {
			call, ok := n.(*ast.CallExpr)
			if !ok {
				return true
			}

			varIds := extractor.ExtractLogVarIds(*call, pass.TypesInfo)
			if len(varIds) == 0 {
				return true
			}

			for _, varId := range varIds {
				for _, pattern := range sensitivePatterns {
					if pattern.MatchString(varId) {
						pass.Reportf(n.Pos(), "Message contains sensitive data: %s", varId)
						break
					}
				}
			}

			return true
		})
	}

	return nil, nil
}
