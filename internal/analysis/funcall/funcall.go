package funcall

import (
	"errors"
	"go/ast"
	"go/token"
	"go/types"
	"strings"
)

var (
	ErrNotLogFunc = errors.New("not log func")
)

// GetLogFuncName returns the name of the logging function if the node is a log call
func GetLogFuncName(
	call ast.CallExpr,
	typeInfo *types.Info,
	logFuncs map[string]struct{},
	pkgName string,
	typeName string,
) (logFuncName string, err error) {
	err = ErrNotLogFunc

	var isLogFuncName = func(s string) bool {
		_, ok := logFuncs[s]
		return ok
	}

	if funcId, ok := call.Fun.(*ast.Ident); ok {
		if !isLogFuncName(funcId.Name) {
			return
		}

		obj := typeInfo.ObjectOf(funcId)
		if fn, ok := obj.(*types.Func); ok {
			return fn.Name(), nil
		}
		return
	}

	sel, ok := call.Fun.(*ast.SelectorExpr)
	if !ok {
		return
	}

	if !isLogFuncName(sel.Sel.Name) {
		return
	}

	logFuncName = sel.Sel.Name

	if pkgId, ok := sel.X.(*ast.Ident); ok {
		obj := typeInfo.ObjectOf(pkgId)
		if obj.(*types.PkgName).Imported().Name() == pkgName {
			return logFuncName, nil
		}
		return
	}

	if typeInfo.TypeOf(sel.X).Underlying().String() == typeName {
		return logFuncName, nil
	}
	return
}

func GetStringArgs(call ast.CallExpr, typeInfo *types.Info) []string {
	var strs []string
	for _, arg := range call.Args {
		switch arg := arg.(type) {
		case *ast.BasicLit:
			if arg.Kind == token.STRING {
				strs = append(strs, strings.Trim(arg.Value, "\""))
			}
		case *ast.Ident:
			if constant, ok := typeInfo.ObjectOf(arg).(*types.Const); ok {
				if val := constant.Val(); val != nil {
					strs = append(strs, strings.Trim(val.String(), "\""))
				}
			}
		}
	}

	return strs
}
