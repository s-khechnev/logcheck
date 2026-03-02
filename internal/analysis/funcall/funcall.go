package funcall

import (
	"go/ast"
	"go/token"
	"go/types"
	"strings"
)

func GetStringArgs(n ast.Node, typeInfo *types.Info) []string {
	call, ok := n.(*ast.CallExpr)
	if !ok {
		panic("unreachable")
	}

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
