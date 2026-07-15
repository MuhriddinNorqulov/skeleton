package file

import (
	"bytes"
	"html/template"

	"example.com/PROJECT_NAME/src/core/domain/ports/file"
)

type HtmlRendererImpl struct{}

// @inject
func NewHtmlRendererImpl() file.HtmlRenderer {
	return &HtmlRendererImpl{}
}

func (this *HtmlRendererImpl) Render(templatePath string, data any) (string, error) {
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", err
	}
	return buf.String(), nil
}
