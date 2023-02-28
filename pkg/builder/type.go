package builder

import "html/template"

type Builder struct {
	IndexTemplate *template.Template
	RepoTemplate  *template.Template
}

func NewBuilder(indexTemplate, repoTemplate *template.Template) *Builder {
	return &Builder{indexTemplate, repoTemplate}
}
