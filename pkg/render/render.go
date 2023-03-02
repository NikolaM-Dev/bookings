package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// RenderTemplate renders a template
func RenderTemplateTest(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl+".page.tmpl", "./templates/base.layout.tmpl")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("err parsing template:", err)
	}
}

var tc = make(map[string]*template.Template)

// RenderTemplate renders a template with cache
func RenderTemplate(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error

	_, inMap := tc[t]
	if !inMap {
		log.Println("creating template and adding to cache")

		err = createTemplateCache(t)
		if err != nil {
			log.Println(err)
		}
	} else {
		log.Println("using cached template")
	}

	tmpl = tc[t]

	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}

// createTemplateCache create a template and add it to the cache
func createTemplateCache(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s.page.tmpl", t),
		"./templates/base.layout.tmpl",
	}

	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}

	tc[t] = tmpl

	return nil
}
