package handlers

import (
	"BBWA/pkg/config"
	"BBWA/pkg/models"
	"BBWA/pkg/render"
	"net/http"
)

//Repo the repository used by the handlers
var Repo *Repository

//Repository is the repository types
type Repository struct {
	App *config.AppConfig
}

//NewRepo Creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

//NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

//Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

//About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	//perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello again"
	//sent the data to the template
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
