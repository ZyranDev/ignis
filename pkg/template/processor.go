package template

import (
	"html/template"
	"os"
	"path/filepath"
	"strings"
)

func ToLower(text string) string {
	return strings.ToLower(text)
}

func ReadTemplate(path string) (tmpl *template.Template, err error) {
	fileContent, err := os.ReadFile(path)
	if err != nil {
		return
	}
	tmpl = template.New(filepath.Base(path)).Funcs(template.FuncMap{
		"ToLower": ToLower,
	})
	return tmpl.Parse(string(fileContent))
}
