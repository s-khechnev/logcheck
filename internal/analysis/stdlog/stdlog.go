package stdlog

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

	funcId, ok := call.Fun.(*ast.Ident) // when . log in imports
	if ok {
		if !isStdLogCall(funcId.Name) {
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

	if !isStdLogCall(sel.Sel.Name) {
		return false
	}

	pkgId, ok := sel.X.(*ast.Ident) // log.Print
	if ok {
		obj := typeInfo.ObjectOf(pkgId)
		if pkgName, ok := obj.(*types.PkgName); ok {
			return pkgName.Imported().Path() == "log"
		} else {
			panic("unreachable")
		}
	}

	return typeInfo.TypeOf(sel.X).Underlying().String() == "*log.Logger"
}

func isStdLogCall(name string) bool {
	stdLogFuncs := map[string]struct{}{
		"Print": {}, "Printf": {}, "Println": {},
		"Fatal": {}, "Fatalf": {}, "Fatalln": {},
		"Panic": {}, "Panicf": {}, "Panicln": {},
	}

	_, ok := stdLogFuncs[name]
	return ok
}

type MessagesExtractor struct {
	stdCallDetector CallDetector
}

func NewMessagesExtractor() *MessagesExtractor {
	return &MessagesExtractor{
		stdCallDetector: CallDetector{},
	}
}

func (e MessagesExtractor) ExtractLogMessages(n ast.Node, typeInfo *types.Info) []string {
	if !e.stdCallDetector.IsLogCall(n, typeInfo) {
		return nil
	}

	return funcall.GetStringArgs(n, typeInfo)
}
