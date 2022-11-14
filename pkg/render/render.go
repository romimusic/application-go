package render

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/romimusic/application-go/pkg/config"
	"github.com/romimusic/application-go/pkg/handlers"
)

// global variables
var functions = template.FuncMap{}
var app *config.AppConfig

// NewTamplates set the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

func RenderTemplate(w http.ResponseWriter, tmpl string, td *handlers.TemplateData) {
	//create a template cache from the app config
	var tc map[string]*template.Template

	if app.UseCache {
		//get the template cache from the app config
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	//get requested template from cache
	t, ok := tc[tmpl]
	if !ok {
		fmt.Println("Could not get template from cache")
	}

	buf := new(bytes.Buffer)

	_ = t.Execute(buf, nil)

	//render template
	_, err := buf.WriteTo(w)

	if err != nil {
		fmt.Println("Error writing template to browser:", err)
	}

}

// CreateTemplateCache creates a template cache as a map
func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.html")

	if err != nil {
		return myCache, err
	}

	//range to the pages
	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.New(name).ParseFiles(page)

		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")

		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts

	}

	return myCache, nil
}

// parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.html")

// err := parsedTemplate.Execute(w, nil)

// if err != nil {
// 	fmt.Println("Error parsing template:", err)
// 	return
// }

// var tc = make(map[string]*template.Template)

// func RenderTemplate(w http.ResponseWriter, t string) {
// 	var tmpl *template.Template
// 	var err error

// 	//check if we already have the template cached
// 	_, inMap := tc[t]

// 	if !inMap {
// 		//if not, parse the template and cache it
// 		err = createTemplateCache(t)

// 		if err != nil {
// 			log.Println("Error creating template cache:", err)
// 		}
// 	}

// 	tmpl = tc[t]

// 	err = tmpl.Execute(w, nil)

// 	if err != nil {
// 		log.Println("Error parsing template:", err)
// 		return
// 	}
// }

// func createTemplateCache(t string) error {
// 	templates := []string{
// 		fmt.Sprintf("./templates/%s", t),
// 		"./templates/base.layout.html",
// 	}

// 	//parse the template file
// 	tmpl, err := template.ParseFiles(templates...)

// 	if err != nil {
// 		return err
// 	}

// 	//add the template to the cache
// 	tc[t] = tmpl

// 	return nil
// }
