package builder

import (
	"bytes"
	"html/template"
	. "ignis/pkg/repository"
	"os"
	"path/filepath"
	"strings"
)

func (b *Builder) Build(outputDir string, repos []Repository) error {
	if err := createDirectories(outputDir); err != nil {
		return err
	}
	context := map[string]interface{}{
		"Repositories": repos,
	}

	/*Build index*/
	if err := WriteTemplate(filepath.Join(outputDir, "index.html"), b.IndexTemplate, context); err != nil {
		return err
	}

	/*Build repositories*/
	repoDir := filepath.Join(outputDir, "repo")
	if err := createDirectories(repoDir); err != nil {
		return err
	}
	for _, repo := range repos {
		err := WriteTemplate(filepath.Join(repoDir, strings.ToLower(repo.Name)+".html"), b.RepoTemplate, repo)
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
