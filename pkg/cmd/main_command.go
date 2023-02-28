package main

import (
	"flag"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"go.zyran.dev/ignis/pkg/builder"
	"go.zyran.dev/ignis/pkg/template"
	"gopkg.in/yaml.v3"
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

	indexTemplate, err := template.ReadTemplate(indexTemplatePath)
	if err != nil {
		log.Info().Err(err)
		return
	}
	repoTemplate, err := template.ReadTemplate(repoTemplatePath)
	if err != nil {
		log.Info().Err(err)
		return
	}

	ignisBuilder := builder.NewBuilder(indexTemplate, repoTemplate)
	config, err := readConfiguration(configPath)
	if config.Repositories == nil {
		log.Info().Msg("No repositories found in the configuration file")
		return
	}
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
