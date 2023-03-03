package config

import (
	"html/template"
	"log"
)

// AppConfig holds the application config
type AppConfig struct {
	InfoLog       *log.Logger
	TemplateCache map[string]*template.Template
	UseCache      bool
}
