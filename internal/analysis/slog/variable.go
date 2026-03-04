package slog

import (
	"go/ast"
	"go/types"
	"logcheck/internal/analysis/funcall"
)

type VarIdsExtractor struct{}

func (e VarIdsExtractor) ExtractLogVarIds(call ast.CallExpr, typeInfo *types.Info) []string {
	if !funcall.IsTargetFuncCall(call, typeInfo, slogLogFuncs, slogPkgName, slogLoggerTypeName) {
		return nil
	}

	var varIds []string
	for _, arg := range call.Args {
		varIds = append(varIds, funcall.ExtractAllIds(arg, typeInfo)...)
	}

	return varIds
}
