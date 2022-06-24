package config

import "html/template"

//AppConfig holds the application config
type AppConfig struct {
	TemplateCash map[string]*template.Template
}
