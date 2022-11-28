package handlers

import (
	"html/template"
	"net/http"
)

var templates *template.Template

func LoadTemplates() {
	templates = template.Must(template.ParseGlob("client/pages/*.html"))
	templates.ParseGlob("client/pages/base/*.html")
}

func Render(w http.ResponseWriter, file string, data interface{}) {
	if err := templates.ExecuteTemplate(w, file, data); err != nil {
		http.Error(w, "There was an exception trying to load the file "+file+"\r\nException detail: "+err.Error(), http.StatusInternalServerError)
	}
}
