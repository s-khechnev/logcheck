package zap

import (
	"go/ast"
	"go/types"

	"github.com/s-khechnev/logcheck/internal/analysis/funcall"
)

type VarIdsExtractor struct{}

func (e VarIdsExtractor) ExtractLogVarIds(call ast.CallExpr, typeInfo *types.Info) []string {
	var isZapLogFunCall = func(logFuncs []string, typName string) bool {
		return funcall.IsTargetFuncCall(call, typeInfo, logFuncs, zapPkgName, typName)
	}

	if !isZapLogFunCall(zapLogFuncs, zapLoggerTypeName) && !isZapLogFunCall(zapSugarLogFuncs, zapSugarLoggerTypeName) {
		return nil
	}

	var varIds []string
	for _, arg := range call.Args {
		varIds = append(varIds, funcall.ExtractAllIds(arg, typeInfo)...)
	}

	return varIds
}
