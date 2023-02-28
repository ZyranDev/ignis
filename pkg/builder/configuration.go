package builder

import "go.zyran.dev/ignis/pkg/repository"

type Configuration struct {
	Host         string                  `yaml:"host"`
	Repositories []repository.Repository `yaml:"repositories"`
}
