// 更新
package sdhelper

import (
	"fmt"
	"runtime"
)

func GetErrorInfo() (file, function string) {
	pc, file, line, _ := runtime.Caller(1)

	file = fmt.Sprintf("error file : %v ( line : %d )", file, line)
	function = runtime.FuncForPC(pc).Name()

	return
}
