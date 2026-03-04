package stdlog

import (
	"go/ast"
	"go/types"

	"github.com/s-khechnev/logcheck/internal/analysis/funcall"
)

var StdLogFuncs = []string{
	"Print", "Printf", "Println",
	"Fatal", "Fatalf", "Fatalln",
	"Panic", "Panicf", "Panicln",
}

type MessagesExtractor struct{}

func (e MessagesExtractor) ExtractLogMessages(call ast.CallExpr, typeInfo *types.Info) []string {
	if !funcall.IsTargetFuncCall(call, typeInfo, StdLogFuncs, "log", "*log.Logger") {
		return nil
	}

	return funcall.ExtractStringArgs(call, typeInfo)
}
