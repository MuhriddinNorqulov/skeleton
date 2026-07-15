package file

type HtmlRenderer interface {
	Render(templatePath string, data any) (string, error)
}
