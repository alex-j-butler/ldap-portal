package helpers

import (
	"gopkg.in/macaron.v1"
	"bytes"
	"text/template"
)

func CheckExists(ctx *macaron.Context, id int) (int) {
	return 404
}

func Format(format string, context map[string]interface{}) string {
	var text bytes.Buffer
	textTemplate := template.New("formatting_template")
	if _, err := textTemplate.Parse(format); err != nil {
		return ""
	}
	if err := textTemplate.Execute(&text, context); err != nil {
		return ""
	}

	return text.String()
}

