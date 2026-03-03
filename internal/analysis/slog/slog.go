package slog

import (
	"go/ast"
	"go/types"
	"logcheck/internal/analysis/funcall"
	"strings"
)

var SlogLogFuncs = map[string]struct{}{
	"Debug": {}, "DebugContext": {},
	"Info": {}, "InfoContext": {},
	"Warn": {}, "WarnContext": {},
	"Error": {}, "ErrorContext": {},
	"Log": {}, "LogContext": {},
	"LogAttrs": {},
}

type MessagesExtractor struct{}

func (e MessagesExtractor) ExtractLogMessages(call ast.CallExpr, typeInfo *types.Info) []string {
	funcName, err := funcall.GetLogFuncName(call, typeInfo, SlogLogFuncs, "slog", "*log/slog.Logger")
	if err != nil || len(call.Args) == 0 {
		return nil
	}

	startIdx := 0
	if strings.HasSuffix(funcName, "Context") {
		startIdx = 1
	}

	var msgs []string

	msg, err := funcall.ExtractStringArg(call.Args[startIdx], typeInfo)
	if err == nil {
		msgs = append(msgs, msg)
	}

	i := startIdx + 1
	for {
		if i >= len(call.Args) {
			break
		}

		keyMsg, err := funcall.ExtractStringArg(call.Args[i], typeInfo)
		if err == nil {
			msgs = append(msgs, keyMsg)
			i += 2
			continue
		}

		if call, ok := call.Args[i].(*ast.CallExpr); ok {
			slogAttrFunc := map[string]struct{}{
				"String": {}, "Int": {},
				"Int64": {}, "Uint64": {},
				"Float64": {}, "Bool": {},
				"Time": {}, "Duration": {},
				"Any": {},
			}
			_, err := funcall.GetLogFuncName(*call, typeInfo, slogAttrFunc, "slog", "*log/slog.Logger")
			if err != nil {
				i += 1
				continue
			}

			keyMsg, err = funcall.ExtractStringArg(call.Args[0], typeInfo)
			if err == nil {
				msgs = append(msgs, keyMsg)
			}

			i += 1
			continue
		}

		i += 1
	}

	return msgs
}
