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

func ReadTemplate(path string, modifiers ...func(*template.FuncMap)) (tmpl *template.Template, err error) {
	fileContent, err := os.ReadFile(path)
	if err != nil {
		return
	}
	funcMap := template.FuncMap{
		"ToLower": ToLower,
	}
	for _, modifier := range modifiers {
		modifier(&funcMap)
	}
	tmpl = template.New(filepath.Base(path)).Funcs(funcMap)
	return tmpl.Parse(string(fileContent))
}
