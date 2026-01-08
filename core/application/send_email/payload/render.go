package payload

import (
	"bytes"
	"html/template"
)

func RenderHTML(payload any, pathHTML string) (string, error) {
	tmpl, err := template.ParseFiles(pathHTML)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, payload); err != nil {
		return "", err
	}

	return buf.String(), nil
}
