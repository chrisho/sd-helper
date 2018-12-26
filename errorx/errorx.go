package errorx

import (
	"github.com/fwhezfwhez/errorx"
)

//获取errorx的错误信息
func GetErrorMessage(v interface{}) string {
	switch t := v.(type) {
	case errorx.Error:
		return v.(errorx.Error).StackTrace()
	case error:
		return v.(error).Error()
	default:
		_ = t
		return ""
	}
}
