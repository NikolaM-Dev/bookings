package render

import (
	"fmt"
	"html/template"
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
