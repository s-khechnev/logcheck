package logcheck

import (
	"go/ast"
	"go/types"
)

type LogMsgExtractor interface {
	ExtractLogMessages(n ast.CallExpr, types *types.Info) []string
}

// LogVarIdsExtractor extracts variable identifiers from log function calls
type LogVarIdsExtractor interface {
	ExtractLogVarIds(n ast.CallExpr, types *types.Info) []string
}
