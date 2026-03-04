package zap

import (
	"errors"
	"go/ast"
	"go/types"
	"logcheck/internal/analysis/funcall"
	"strings"
)

var (
	ErrNotZapField = errors.New("not zap field")
)

var (
	zapLogFuncs = []string{
		"Debug", "DPanic",
		"Info", "Error",
		"Warn", "Panic",
		"Fatal", "Log",
	}

	zapSugarLogFuncs = []string{
		"Debug", "Info", "Warn", "Error", "DPanic", "Panic", "Fatal",
		"Debugf", "Infof", "Warnf", "Errorf", "DPanicf", "Panicf", "Fatalf",
		"Debugw", "Infow", "Warnw", "Errorw", "DPanicw", "Panicw", "Fatalw",
		"Debugln", "Infoln", "Warnln", "Errorln", "DPanicln", "Panicln", "Fatalln",
	}

	zapFieldFuncs = []string{
		"Any", "Array", "Binary", "Bool", "Bools",
		"ByteString", "ByteStrings", "Complex128", "Complex128s", "Complex64",
		"Complex64s", "Duration", "Durations", "Error", "Errors",
		"Float32", "Float32s", "Float64", "Float64s", "Inline",
		"Int", "Ints", "Int8", "Int8s", "Int16",
		"Int16s", "Int32", "Int32s", "Int64", "Int64s",
		"Namespace", "Object", "Reflect", "Stack",
		"StackSkip", "String", "Strings", "Stringer", "Time",
		"Times", "Timep", "Uint", "Uints", "Uint8",
		"Uint8s", "Uint16", "Uint16s", "Uint32", "Uint32s",
		"Uint64", "Uint64s", "Uintptr", "Uintptrs",
	}
)

const (
	zapPkgName             = "zap"
	zapLoggerTypeName      = "*go.uber.org/zap.Logger"
	zapSugarLoggerTypeName = "*go.uber.org/zap.SugaredLogger"
)

type MessagesExtractor struct{}

func (e MessagesExtractor) ExtractLogMessages(call ast.CallExpr, typeInfo *types.Info) []string {
	switch {
	case funcall.IsTargetFuncCall(call, typeInfo, zapLogFuncs, zapPkgName, zapLoggerTypeName):
		return ExtractMsgsFromStructuredLog(call, typeInfo)
	case funcall.IsTargetFuncCall(call, typeInfo, zapSugarLogFuncs, zapPkgName, zapSugarLoggerTypeName):
		s, err := funcall.ExtractFuncName(call.Fun)
		if err != nil {
			return nil
		}

		if strings.HasSuffix(s, "w") {
			return ExtractMsgsFromStructuredLog(call, typeInfo)
		}
		return funcall.ExtractStringArgs(call, typeInfo)
	}

	return nil
}

func ExtractMsgsFromStructuredLog(call ast.CallExpr, typeInfo *types.Info) []string {
	var msgs []string

	startIdx := 0
	for i, arg := range call.Args {
		msg, err := funcall.ExtractString(arg, typeInfo)
		if err == nil {
			msgs = append(msgs, msg)
			startIdx = i
			break
		}
	}

	i := startIdx + 1
	for {
		if i >= len(call.Args) {
			break
		}

		// logger.Info("msg", "keyMsg", "val1")
		keyMsg, err := funcall.ExtractString(call.Args[i], typeInfo)
		if err == nil {
			msgs = append(msgs, keyMsg)
			i += 2
			continue
		}

		// logger.Info("msg", logger.String("keyMsg", "val1"))
		keyMsg, err = ExtractKeyFromZapField(call.Args[i], typeInfo)
		if err == nil {
			msgs = append(msgs, keyMsg)
			i += 1
			continue
		}

		//	log.Infow("hel",
		//		zap.Dict("Први",
		//			zap.Dict("Ключ1", zap.String("Ключ2", "val"))
		//			))
		//  dictMsgs = ["Први", "Ключ1", "Ключ2"]
		dictMsgs := ExtractMsgsFromZapDict(call.Args[i], typeInfo)
		msgs = append(msgs, dictMsgs...)

		i += 1
	}

	return msgs
}

func ExtractMsgsFromZapDict(call ast.Expr, typeInfo *types.Info) []string {
	var result []string

	var helper func(call ast.Expr, typeInfo *types.Info)
	helper = func(call ast.Expr, typeInfo *types.Info) {
		if call, ok := call.(*ast.CallExpr); ok {
			if !funcall.IsTargetFuncCall(*call, typeInfo, []string{"Dict"}, zapPkgName, zapLoggerTypeName) {
				return
			}

			groupKeyMsg, err := funcall.ExtractString(call.Args[0], typeInfo)
			if err != nil {
				return
			}

			result = append(result, groupKeyMsg)

			if len(call.Args) == 1 {
				return
			}

			for _, arg := range call.Args[1:] {
				keyMsg, err := ExtractKeyFromZapField(arg, typeInfo)
				if err != nil {
					helper(arg, typeInfo)
				} else {
					result = append(result, keyMsg)
				}
			}
		}
	}

	helper(call, typeInfo)

	return result
}

func ExtractKeyFromZapField(call ast.Expr, typeInfo *types.Info) (string, error) {
	if call, ok := call.(*ast.CallExpr); ok {
		if !funcall.IsTargetFuncCall(*call, typeInfo, zapFieldFuncs, zapPkgName, zapLoggerTypeName) {
			return "", ErrNotZapField
		}

		keyMsg, err := funcall.ExtractString(call.Args[0], typeInfo)
		if err == nil {
			return keyMsg, nil
		}
	}

	return "", ErrNotZapField
}
