package utils

import (
	"bytes"
	"github.com/kakuilan/kgo"
	"strings"
	"text/template"
)

func UnicodeIndex(str, substr string) int {
	// 子串在字符串的字节位置
	result := strings.Index(str, substr)
	if result >= 0 {
		// 获得子串之前的字符串并转换成[]byte
		prefix := []byte(str)[0:result]
		// 将子串之前的字符串转换成[]rune
		rs := []rune(string(prefix))
		// 获得子串之前的字符串的长度，便是子串在字符串的字符位置
		result = len(rs)
	}

	return result
}

func ReplaceString(pattern, replace, src string) (string, error) {
	if r, err := GetRegexp(pattern); err == nil {
		return r.ReplaceAllString(src, replace), nil
	} else {
		return "", err
	}
}

func Contains(haystack, needle string, startOffset ...int) int {
	length := len(haystack)
	offset := 0
	if len(startOffset) > 0 {
		offset = startOffset[0]
	}
	if length == 0 || offset > length || -offset > length {
		return -1
	}

	if offset < 0 {
		offset += length
	}
	pos := strings.Index(strings.ToLower(haystack[offset:]), strings.ToLower(needle))
	if pos == -1 {
		return -1
	}
	return pos + offset
}

// 字符串模板解析
func TemplateResolve(temp string, data any) string {
	t, _ := template.New("string-temp").Parse(temp)
	var tmplBytes bytes.Buffer

	err := t.Execute(&tmplBytes, data)
	if err != nil {
		panic(err)
	}
	return tmplBytes.String()
}

func ReverStrTemplate(temp, str string, res map[string]any) {
	index := UnicodeIndex(temp, "{")
	ei := UnicodeIndex(temp, "}") + 1
	next := kgo.KStr.Trim(temp[ei:], " ")
	nextContain := UnicodeIndex(next, "{")
	nextIndexValue := next
	if nextContain != -1 {
		nextIndexValue = kgo.KStr.Substr(next, 0, nextContain)
	}
	key := temp[index+1 : ei-1]
	// 如果后面没有内容了，则取字符串的长度即可
	var valueLastIndex int
	if nextIndexValue == "" {
		valueLastIndex = kgo.KStr.MbStrlen(str)
	} else {
		valueLastIndex = UnicodeIndex(str, nextIndexValue)
	}
	value := kgo.KStr.Trim(kgo.KStr.Substr(str, index, valueLastIndex), " ")
	res[key] = value
	// 如果后面的还有需要解析的，则递归调用解析
	if nextContain != -1 {
		ReverStrTemplate(next, kgo.KStr.Trim(kgo.KStr.Substr(str, UnicodeIndex(str, value)+kgo.KStr.MbStrlen(value), kgo.KStr.MbStrlen(str))), res)
	}
}

func B2S(bs []uint8) string {
	ba := []byte{}
	for _, b := range bs {
		ba = append(ba, byte(b))
	}
	return string(ba)
}

func IdsStrToIdsIntGroup(keys string) []int64 {
	IDS := make([]int64, 0)
	ids := strings.Split(keys, ",")
	for i := 0; i < len(ids); i++ {
		ID := kgo.KConv.Str2Int(ids[i])
		IDS = append(IDS, int64(ID))
	}
	return IDS
}

// 获取部门
// isP 是父ID 否则子ID
func DeptPCIds(deptIds []string, id int64, isP bool) []int64 {
	pRes := make([]int64, 0)
	cRes := make([]int64, 0)
	is := true
	for _, deptId := range deptIds {
		did := kgo.KConv.Str2Int64(deptId)
		if is {
			pRes = append(pRes, did)
		}
		if kgo.KConv.Str2Int64(deptId) == id {
			is = false
		}
		if !is {
			cRes = append(cRes, did)
		}
	}
	if isP {
		return pRes
	} else {
		return cRes
	}
}
