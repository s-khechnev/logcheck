package stdlog

import (
	"go/ast"
	"go/types"
	"logcheck/internal/analysis/funcall"
)

var StdLogFuncs = map[string]struct{}{
	"Print": {}, "Printf": {}, "Println": {},
	"Fatal": {}, "Fatalf": {}, "Fatalln": {},
	"Panic": {}, "Panicf": {}, "Panicln": {},
}

type MessagesExtractor struct{}

func (e MessagesExtractor) ExtractLogMessages(call ast.CallExpr, typeInfo *types.Info) []string {
	_, err := funcall.GetLogFuncName(call, typeInfo, StdLogFuncs, "log", "*log.Logger")
	if err != nil {
		return nil
	}

	return funcall.GetStringArgs(call, typeInfo)
}
