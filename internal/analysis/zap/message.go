package zap

import (
	"errors"
	"go/ast"
	"go/types"
	"logcheck/internal/analysis/funcall"
)

var (
	zapLogFuncs = []string{
		"Debug", "DebugContext",
		"Info", "InfoContext",
		"Warn", "WarnContext",
		"Error", "ErrorContext",
		"Log", "LogContext",
		"LogAttrs",
	}
	slogAttrFunc = []string{
		"String", "Int",
		"Int64", "Uint64",
		"Float64", "Bool",
		"Time", "Duration",
		"Any",
	}
)

const (
	zapPkgName        = "go.uber.org/zap"
	zapLoggerTypeName = "*go.uber.org/zap.Logger\n"
)

type MessagesExtractor struct{}

func (e MessagesExtractor) ExtractLogMessages(call ast.CallExpr, typeInfo *types.Info) []string {
	if !funcall.IsTargetFuncCall(call, typeInfo, zapLogFuncs, zapPkgName, zapLoggerTypeName) {
		return nil
	}

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

		// slog.Info("msg", "keyMsg", "val1")
		keyMsg, err := funcall.ExtractString(call.Args[i], typeInfo)
		if err == nil {
			msgs = append(msgs, keyMsg)
			i += 2
			continue
		}

		// slog.Info("msg", slog.String("keyMsg", "val1"))
		keyMsg, err = ExtractKeyFromSlogAttr(call.Args[i], typeInfo)
		if err == nil {
			msgs = append(msgs, keyMsg)
			i += 1
			continue
		}

		//slog.Info("Aboba", slog.Group("request",
		//	slog.Group("xyy", slog.String("method", "GET"),
		//		slog.Int("status", 200)),
		//	slog.String("qwe1", "val"),
		//	slog.String("qwe2", "val"),
		//))
		// groupMsgs will be [request xyy method status qwe1 qwe2]
		groupMsgs := ExtractMsgsFromSlogGroup(call.Args[i], typeInfo)
		msgs = append(msgs, groupMsgs...)

		i += 1
	}

	return msgs
}

func ExtractMsgsFromSlogGroup(call ast.Expr, typeInfo *types.Info) []string {
	var result []string

	var helper func(call ast.Expr, typeInfo *types.Info)
	helper = func(call ast.Expr, typeInfo *types.Info) {
		if call, ok := call.(*ast.CallExpr); ok {
			if !funcall.IsTargetFuncCall(*call, typeInfo, []string{"Group"}, zapPkgName, zapLoggerTypeName) {
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
				keyMsg, err := ExtractKeyFromSlogAttr(arg, typeInfo)
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

var (
	ErrNotSlogAttr = errors.New("not slog attr")
)

func ExtractKeyFromSlogAttr(call ast.Expr, typeInfo *types.Info) (string, error) {
	if call, ok := call.(*ast.CallExpr); ok {
		if !funcall.IsTargetFuncCall(*call, typeInfo, slogAttrFunc, zapPkgName, zapLoggerTypeName) {
			return "", ErrNotSlogAttr
		}

		keyMsg, err := funcall.ExtractString(call.Args[0], typeInfo)
		if err == nil {
			return keyMsg, nil
		}
	}

	return "", ErrNotSlogAttr
}
