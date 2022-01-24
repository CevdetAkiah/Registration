package config

import "html/template"

type AppConfig struct {
	InProduction  bool
	UseCache      bool
	TemplateCache map[string]*template.Template
}
