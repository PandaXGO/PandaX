package tool

import (
	"reflect"
	"regexp"
	"strings"
)

func ToCamelCase(s string) string {
	re := regexp.MustCompile(`[_\W]+`)
	words := re.Split(s, -1)
	for i := range words {
		if i != 0 {
			words[i] = strings.Title(words[i])
		}
	}
	return strings.Join(words, "")
}

func RegexpKey(str string) []string {
	// 定义正则表达式
	re := regexp.MustCompile(`\${([^}]+)}`)
	matches := re.FindAllStringSubmatch(str, -1)
	// 提取匹配项的内容
	var results []string
	for _, match := range matches {
		if len(match) >= 2 {
			results = append(results, match[1])
		}
	}
	return results
}

func RegexpGetSql(str string) string {
	// 定义正则表达式
	re := regexp.MustCompile(`\${([^}]+)}`)
	return re.ReplaceAllString(str, "?")
}

func GetStructKeys(obj interface{}) []string {
	val := reflect.ValueOf(obj)
	typ := val.Type()

	keys := make([]string, 0, typ.NumField())

	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		keys = append(keys, field.Name)
	}

	return keys
}

func GetMapKeys(obj map[string]interface{}) []string {
	keys := make([]string, 0, len(obj))

	for key := range obj {
		keys = append(keys, key)
	}

	return keys
}

func CheckInterfaceIsArray(data interface{}) (bool, []map[string]interface{}) {
	if data == nil {
		return false, nil
	}
	valueType := reflect.TypeOf(data)
	// 判断类型是否为数组或切片
	if valueType.Kind() == reflect.Slice || valueType.Kind() == reflect.Array {
		var maps []map[string]interface{}
		for _, item := range data.([]interface{}) {
			if m, ok := item.(map[string]interface{}); ok {
				maps = append(maps, m)
			}
		}
		return true, maps
	}
	return false, nil
}

func GetInterfaceType(v interface{}) string {
	interfaceType := reflect.TypeOf(v)
	return interfaceType.String()
}
