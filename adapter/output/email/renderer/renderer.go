package renderer

import (
	"bytes"
	"errors"
	"html/template"
)

// HTMLEmailRenderer is responsible for rendering email bodies from HTML templates.
//
// It resolves a template path based on a template identifier and executes
// the template using the provided data, returning the rendered HTML as a string.
type HTMLEmailRenderer struct {
}

// Render renders an HTML email template identified by templateID using the given data.
//
// It looks up the template path, parses the HTML file, and executes it.
// An error is returned if the template does not exist or if rendering fails.
func (r *HTMLEmailRenderer) Render(templateID string, data any) (string, error) {
	path, ok := pathTemplates[templateID]
	if !ok {
		return "", errors.New("template not found")
	}

	return r.renderHTML(data, path)
}

// renderHTML loads and executes the HTML template located at pathHTML
// using the provided data, returning the rendered output as a string.
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
