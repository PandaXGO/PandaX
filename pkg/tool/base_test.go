package tool

import "testing"

func TestToCamelCase(t *testing.T) {
	camelCase := ToCamelCase("hello_world")
	t.Log(camelCase)
}
