package logcheck

import (
	"go/ast"
	"go/types"
)

type LogMsgExtractor interface {
	ExtractLogMessages(n ast.CallExpr, types *types.Info) []string
}
