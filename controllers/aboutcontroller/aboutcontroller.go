package aboutcontroller

import (
	"html/template"
	"net/http"
)
var tmpl *template.Template

func init(){
	tmpl = template.Must(template.ParseGlob("views/home/*.html"))
}

func Index(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "about.html", nil)
}
