package logcheck

import (
	"go/ast"
	"go/types"
)

type LogCallDetector interface {
	IsLogCall(n ast.Node, types *types.Info) bool
}

type LogMsgExtractor interface {
	ExtractLogMessages(n ast.Node, types *types.Info) []string
}
