// 更新
package sdhelper

import (
	"os"
)

func GetEnv(key string) (value string) {
	return os.Getenv(key)
}

func SetEnv(key string, value string) error {
	return os.Setenv(key, value)
}