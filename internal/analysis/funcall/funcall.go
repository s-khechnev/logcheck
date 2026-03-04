package funcall

import (
	"fmt"
	"go/ast"
	"go/token"
	"go/types"
	"slices"
	"strings"
)

func ExtractString(expr ast.Expr, typeInfo *types.Info) (string, error) {
	var stringBuilder strings.Builder

	ast.Inspect(expr, func(n ast.Node) bool {
		switch n := n.(type) {
		case *ast.BinaryExpr:
			return n.Op == token.ADD
		case *ast.BasicLit:
			if n.Kind == token.STRING {
				stringBuilder.WriteString(strings.Trim(n.Value, "\""))
			}
		case *ast.Ident:
			if constant, ok := typeInfo.ObjectOf(n).(*types.Const); ok {
				if val := constant.Val(); val != nil {
					stringBuilder.WriteString(strings.Trim(val.String(), "\""))
				}
			}
		}

		return false
	})

	if stringBuilder.Len() == 0 {
		return "", fmt.Errorf("not string arg")
	}

	return stringBuilder.String(), nil
}

func ExtractStringArgs(call ast.CallExpr, typeInfo *types.Info) []string {
	var strs []string
	for _, arg := range call.Args {
		strArg, err := ExtractString(arg, typeInfo)
		if err != nil {
			continue
		}
		strs = append(strs, strArg)
	}

	return strs
}

// IsTargetFuncCall determines whether the given function call is a call to a target function.
// The function supports three verification scenarios:
//  1. Direct function call: name() - checks if name is in targetFuncNames
//  2. Package function call: pkg.Name() - checks function name and package name
//  3. Method call: var.Name() - checks function name and receiver type
//
// Parameters:
//   - call: AST node representing the function call
//   - typeInfo: type information obtained from types.Info
//   - targetFuncNames: list of function names considered as targets
//   - pkgName: package name to verify for calls like pkg.Func()
//   - typeName: string representation of the type to verify for method calls
func IsTargetFuncCall(
	call ast.CallExpr,
	typeInfo *types.Info,
	targetFuncNames []string,
	pkgName string,
	typeName string,
) bool {
	var isTargetFuncName = func(s string) bool {
		return slices.Contains(targetFuncNames, s)
	}

	if funcId, ok := call.Fun.(*ast.Ident); ok {
		return isTargetFuncName(funcId.Name)
	}

	sel, ok := call.Fun.(*ast.SelectorExpr)
	if !ok {
		return false
	}

	if !isTargetFuncName(sel.Sel.Name) {
		return false
	}

	if pkgId, ok := sel.X.(*ast.Ident); ok {
		obj := typeInfo.ObjectOf(pkgId)

		if pkg, ok := obj.(*types.PkgName); ok && pkg.Imported().Name() == pkgName {
			return true
		}
	}

	return typeInfo.TypeOf(sel.X).Underlying().String() == typeName
}

func ExtractAllIds(expr ast.Expr, typeInfo *types.Info) []string {
	var varIds []string

	ast.Inspect(expr, func(n ast.Node) bool {
		switch n := n.(type) {
		case *ast.Ident:
			if obj := typeInfo.ObjectOf(n); obj != nil {
				switch obj.(type) {
				case *types.Var, *types.Const:
					varIds = append(varIds, n.Name)
				}
			}
		case *ast.SelectorExpr:
			if sel, ok := typeInfo.Selections[n]; ok && sel != nil {
				if obj := sel.Obj(); obj != nil {
					if _, ok := obj.(*types.Var); ok {
						varIds = append(varIds, n.Sel.Name)
					}
				}
			}
		}

		return true
	})

	return varIds
}
