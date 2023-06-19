package tool

import (
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
