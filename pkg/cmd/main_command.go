package main

import (
	"flag"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"go.zyran.dev/ignis/pkg/builder"
	"go.zyran.dev/ignis/pkg/repository"
	"go.zyran.dev/ignis/pkg/template"
	"gopkg.in/yaml.v3"
	template2 "html/template"
	"os"
)

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
}

func main() {
	var configPath, indexTemplatePath, repoTemplatePath, outputDir string

	flag.StringVar(&configPath, "config", "config.yml", "Configuration file path")
	flag.StringVar(&indexTemplatePath, "index", "templates/index.html", "Index template path")
	flag.StringVar(&repoTemplatePath, "repo", "templates/repository.html", "Repository template path")
	flag.StringVar(&outputDir, "output", "build", "Output directory path")
	flag.Parse()

	config, err := readConfiguration(configPath)
	if config.Repositories == nil {
		log.Info().Msg("No repositories found in the configuration file")
		return
	}

	getHostFunc := func(funcMap *template2.FuncMap) {
		(*funcMap)["GetHost"] = func(repo repository.Repository) string {
			return config.Host
		}
	}

	indexTemplate, err := template.ReadTemplate(indexTemplatePath, getHostFunc)
	if err != nil {
		log.Info().Err(err)
		return
	}
	repoTemplate, err := template.ReadTemplate(repoTemplatePath, getHostFunc)
	if err != nil {
		log.Info().Err(err)
		return
	}

	ignisBuilder := builder.NewBuilder(indexTemplate, repoTemplate)

	err = ignisBuilder.Build(outputDir, config)
	if err != nil {
		log.Info().Err(err)
		return
	}
	log.Info().Msg("Successfully generated ignis files.")
}

func readConfiguration(path string) (config *builder.Configuration, err error) {
	config = &builder.Configuration{}
	fileContent, err := os.ReadFile(path)
	if err != nil {
		return
	}
	err = yaml.Unmarshal(fileContent, config)
	return
}
