package hypertext

import (
	"embed"
	"html/template"
)

//go:embed templates/*.gohtml
var templateFiles embed.FS

var templates = template.Must(template.ParseFS(templateFiles, "templates/*.gohtml"))
