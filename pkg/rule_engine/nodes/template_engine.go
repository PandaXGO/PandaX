package nodes

import (
	"bytes"
	"text/template"
)

func ParseTemplate(content string, data map[string]interface{}) (string, error) {
	tmpl, err := template.New("template").Parse(content)
	if err != nil {
		return "", err
	}
	buffer := &bytes.Buffer{}
	err = tmpl.Execute(buffer, data)
	if err != nil {
		return "", err
	}
	return buffer.String(), nil
}
