package stringx

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// Trim string space
func TrimStringSpace(str string) string {
	return strings.Trim(str, " ")
}

// Trim string cutset
func TrimStringCutset(str, cutset string) string {
	return strings.Trim(str, cutset)
}

// string to int64
func String2Int64(str string) (int64, error) {
	return strconv.ParseInt(str, 10, 64)
}

// string to int32
func String2Int32(str string) int32 {
	int64Str, _ := strconv.Atoi(str)

	return int32(int64Str)
}

func Float2String(number float64, bitSize int) string {
	return strconv.FormatFloat(number, 'f', -1, bitSize)
}

// map map[int]string
// map 的 key 必须从 0 开始，顺序排列
// sep is placed between elements in the resulting string.
func ImplodeMapIntString(maps map[int]string, sep string) (str string) {
	mapLen := len(maps)

	for i := 0; i < mapLen; i++ {
		str += maps[i] + sep
	}

	str = strings.TrimRight(str, sep)

	return
}

// map map[int32]string
// map 的 key 必须从 0 开始，顺序排列
// sep is placed between elements in the resulting string.
func ImplodeMapInt32String(maps map[int32]string, sep string) (str string) {
	mapLen := int32(len(maps))

	for i := int32(0); i < mapLen; i++ {
		str += maps[i] + sep
	}

	str = strings.TrimRight(str, sep)

	return
}

// 将"_"转换为"-"
func ConvertUnderlineToWhippletree(str string) string {
	return strings.Replace(str, "_", "-", -1)
}

// snake string, XxYy to xx_yy , XxYY to xx_yy
func SnakeString(s string) string {
	data := make([]byte, 0, len(s)*2)
	j := false
	num := len(s)
	for i := 0; i < num; i++ {
		d := s[i]
		if i > 0 && d >= 'A' && d <= 'Z' && j {
			data = append(data, '_')
		}
		if d != '_' {
			j = true
		}
		data = append(data, d)
	}
	return strings.ToLower(string(data[:]))
}

// camel string, xx_yy to XxYy
func CamelString(s string) string {
	data := make([]byte, 0, len(s))
	j := false
	k := false
	num := len(s) - 1
	for i := 0; i <= num; i++ {
		d := s[i]
		if k == false && d >= 'A' && d <= 'Z' {
			k = true
		}
		if d >= 'a' && d <= 'z' && (j || k == false) {
			d = d - 32
			j = false
			k = true
		}
		if k && d == '_' && num > i && s[i+1] >= 'a' && s[i+1] <= 'z' {
			j = true
			continue
		}
		data = append(data, d)
	}
	return string(data[:])
}

var searchHandle = regexp.MustCompile(`{\S+}`)

//字符串命名参数格式化，明明参数用 { } 包裹，类似于python “xxx{data}xxx”.format(data=aaa)
func NameFormat(formatStr string, data map[string]string) (string, error) {
	length := len(data)
	if length == 0 {
		return formatStr, nil
	}
	var err error
	result := searchHandle.ReplaceAllStringFunc(formatStr, func(paramName string) string {
		if strings.Contains(paramName, "{{") && strings.Contains(paramName, "}}") {
			return paramName
		}
		paramName = strings.TrimRight(strings.TrimLeft(paramName, "{"), "}")
		val, ok := data[paramName]
		if !ok {
			err = fmt.Errorf("%s not found", paramName)
			return ""
		}
		return val
	})
	if nil != err {
		return "", err
	}
	return result, nil
}
