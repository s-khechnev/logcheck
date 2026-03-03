package funcall

import (
	"errors"
	"fmt"
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

func ExtractStringArg(expr ast.Expr, typeInfo *types.Info) (string, error) {
	switch expr := expr.(type) {
	case *ast.BasicLit:
		if expr.Kind == token.STRING {
			return strings.Trim(expr.Value, "\""), nil
		}
	case *ast.Ident:
		if constant, ok := typeInfo.ObjectOf(expr).(*types.Const); ok {
			if val := constant.Val(); val != nil {
				return strings.Trim(val.String(), "\""), nil
			}
		}
	}

	return "", fmt.Errorf("not string arg")
}

func GetStringArgs(call ast.CallExpr, typeInfo *types.Info) []string {
	var strs []string
	for _, arg := range call.Args {
		strArg, err := ExtractStringArg(arg, typeInfo)
		if err != nil {
			continue
		}
		strs = append(strs, strArg)
	}

	return strs
}

//func F(
//	call ast.CallExpr,
//	typeInfo *types.Info,
//	funcs map[string]struct{},
//	pkgName string,
//	typeName string,
//) bool {
//	var isNeededFunc = func(s string) bool {
//		_, ok := funcs[s]
//		return ok
//	}
//
//	if funcId, ok := call.Fun.(*ast.Ident); ok {
//		obj := typeInfo.ObjectOf(funcId)
//		if fn, ok := obj.(*types.Func); ok {
//			return isNeededFunc(fn.Name())
//		}
//		return false
//	}
//
//	sel, ok := call.Fun.(*ast.SelectorExpr)
//	if !ok {
//		return false
//	}
//
//	if !isNeededFunc(sel.Sel.Name) {
//		return false
//	}
//
//	if pkgId, ok := sel.X.(*ast.Ident); ok {
//		obj := typeInfo.ObjectOf(pkgId)
//		return obj.(*types.PkgName).Imported().Name() == pkgName
//	}
//
//	return typeInfo.TypeOf(sel.X).Underlying().String() == typeName
//}
