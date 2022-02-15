package config

import (
	"html/template"

	"github.com/CevdetAkiah/Registration/pkg/models"
	"github.com/alexedwards/scs/v2"
)

type AppConfig struct {
	InProduction  bool
	UseCache      bool
	TemplateCache map[string]*template.Template
	Session       *scs.SessionManager
	DbUsers       map[string]models.User
	DbSessions    map[string]string
}
