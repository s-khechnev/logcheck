package slog

import (
	"go/ast"
	"go/types"
	"logcheck/internal/analysis/funcall"
)

var SlogLogFuncs = map[string]struct{}{
	"Debug": {}, "DebugContext": {},
	"Info": {}, "InfoContext": {},
	"Warn": {}, "WarnContext": {},
	"Error": {}, "ErrorContext": {},
	"Log": {}, "LogContext": {},
}

type MessagesExtractor struct{}

func (e MessagesExtractor) ExtractLogMessages(n ast.CallExpr, typeInfo *types.Info) []string {
	_, err := funcall.GetLogFuncName(n, typeInfo, SlogLogFuncs, "slog", "*log/slog.Logger")
	if err != nil {
		return nil
	}

	return funcall.GetStringArgs(n, typeInfo)
}
