package handlers

import (
	"net/http"

	"github.com/NikolaM-Dev/bookings/pkg/config"
	"github.com/NikolaM-Dev/bookings/pkg/models"
	"github.com/NikolaM-Dev/bookings/pkg/render"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home", &models.TemplateData{})
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// perfom some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again"

	// send the data to the template
	render.RenderTemplate(w, "about", &models.TemplateData{
		StringMap: stringMap,
	})
}
