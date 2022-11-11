package handlers

import (
	"net/http"

	"github.com/romimusic/application-go/pkg/render"
)

func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.html")
}

// About is the about page
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.html")
}
