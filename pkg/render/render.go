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
