package utils

import "runtime"

func GetCallerFuncName(skip int) string {
	if skip <= 2 {
		skip = 2
	}
	pc, _, _, _ := runtime.Caller(skip)
	function := runtime.FuncForPC(pc)
	if function == nil {
		return ""
	}
	return function.Name()
}
