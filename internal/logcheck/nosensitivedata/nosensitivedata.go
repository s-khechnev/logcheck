package nosensitivedata

import (
	"go/ast"
	"regexp"

	"github.com/s-khechnev/logcheck/internal/logcheck"

	"golang.org/x/tools/go/analysis"
)

var defaultSensitivePatterns = []string{
	`(?i)pass(word)?|pwd`,
	`(?i)token|jwt|refresh|access_token`,
	`(?i)api[_-]?key|apikey|api_secret`,
	`(?i)secret|private_key`,
	`(?i)card|ccv|cvv|credit_card`,
	`(?i)auth|authorization|basic_auth`,
	`(?i)session|cookie|session_id`,
	`(?i)conn(ection)?[_-]?string|dsn`,
	`(?i)email|e-?mail|mail`,
}

func NewAnalyzer(loggerName string, extractor logcheck.LogVarIdsExtractor, patterns ...string) *analysis.Analyzer {
	if len(patterns) == 0 {
		patterns = defaultSensitivePatterns
	}

	regexps := make([]*regexp.Regexp, len(patterns))
	for i, pattern := range patterns {
		regexps[i] = regexp.MustCompile(pattern)
	}

	return &analysis.Analyzer{
		Name: loggerName + "_nosensetivedata",
		Doc:  "Check that the log message doesn't contains sensitive data",
		Run: func(pass *analysis.Pass) (any, error) {
			return run(extractor, pass, regexps)
		},
	}
}

func run(extractor logcheck.LogVarIdsExtractor, pass *analysis.Pass, patterns []*regexp.Regexp) (any, error) {
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
				for _, pattern := range patterns {
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
