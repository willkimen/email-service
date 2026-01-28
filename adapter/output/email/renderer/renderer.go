package renderer

import (
	"bytes"
	"errors"
	"html/template"
)

type HTMLEmailRenderer struct {
}

func (r *HTMLEmailRenderer) Render(templateID string, data any) (string, error) {
	path, ok := pathTemplates[templateID]
	if !ok {
		return "", errors.New("template not found")
	}

	return r.renderHTML(data, path)
}

func (r *HTMLEmailRenderer) renderHTML(data any, pathHTML string) (string, error) {
	tmpl, err := template.ParseFiles(pathHTML)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}
