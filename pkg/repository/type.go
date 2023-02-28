package repository

type Repository struct {
	Name     string `yaml:"name"`
	VSC      string `yaml:"vsc"`
	RepoURL  string `yaml:"repo"`
	Branch   string `yaml:"branch"`
	Redirect string `yaml:"redirect"`
}
