package builder

import (
	"bytes"
	"html/template"
	"os"
	"path/filepath"
	"strings"
)

func (b *Builder) Build(outputDir string, config *Configuration) error {
	if err := createDirectories(outputDir); err != nil {
		return err
	}

	/*Build index*/
	if err := WriteTemplate(filepath.Join(outputDir, "index.html"), b.IndexTemplate, config); err != nil {
		return err
	}

	/*Build repositories*/
	for _, repo := range config.Repositories {
		err := WriteTemplate(filepath.Join(outputDir, strings.ToLower(repo.Name)+".html"), b.RepoTemplate, repo)
		if err != nil {
			return err
		}
	}
	return nil
}

func createDirectories(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}

func ExecuteTemplate(template *template.Template, data interface{}) (writer bytes.Buffer, err error) {
	err = template.Execute(&writer, data)
	return
}

func WriteTemplate(output string, template *template.Template, data interface{}) error {
	buffer, err := ExecuteTemplate(template, data)
	if err != nil {
		return err
	}
	return os.WriteFile(output, buffer.Bytes(), os.ModePerm)
}
