package slog

import (
	"go/ast"
	"go/types"
	"logcheck/internal/analysis/funcall"
)

type CallDetector struct{}

func (CallDetector) IsLogCall(n ast.Node, typeInfo *types.Info) bool {
	call, ok := n.(*ast.CallExpr)
	if !ok {
		return false
	}

	funcId, ok := call.Fun.(*ast.Ident) // when . slog in imports
	if ok {
		if !isSlogCall(funcId.Name) {
			return false
		}

		obj := typeInfo.ObjectOf(funcId)
		if fn, ok := obj.(*types.Func); ok {
			pkg := fn.Pkg()
			if pkg == nil {
				return false
			}

			return pkg.Path() == "log"
		}
	}

	sel, ok := call.Fun.(*ast.SelectorExpr)
	if !ok {
		return false
	}

	if !isSlogCall(sel.Sel.Name) {
		return false
	}

	pkgId, ok := sel.X.(*ast.Ident) // slog.Print
	if ok {
		obj := typeInfo.ObjectOf(pkgId)
		if pkgName, ok := obj.(*types.PkgName); ok {
			return pkgName.Imported().Path() == "log/slog"
		} else {
			panic("unreachable")
		}
	}

	return typeInfo.TypeOf(sel.X).Underlying().String() == "*log/slog.Logger"
}

func isSlogCall(name string) bool {
	slogFuncs := map[string]struct{}{
		"Debug": {}, "DebugContext": {},
		"Info": {}, "InfoContext": {},
		"Warn": {}, "WarnContext": {},
		"Error": {}, "ErrorContext": {},
		"Log": {}, "LogContext": {},
	}

	_, ok := slogFuncs[name]
	return ok
}

type MessagesExtractor struct {
	slogCallDetector CallDetector
}

func NewMessagesExtractor() *MessagesExtractor {
	return &MessagesExtractor{
		slogCallDetector: CallDetector{},
	}
}

func (e MessagesExtractor) ExtractLogMessages(n ast.Node, typeInfo *types.Info) []string {
	if !e.slogCallDetector.IsLogCall(n, typeInfo) {
		return nil
	}

	return funcall.GetStringArgs(n, typeInfo)
}
