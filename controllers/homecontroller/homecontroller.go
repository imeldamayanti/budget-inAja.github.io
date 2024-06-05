package homecontroller

import (
	"html/template"
	"net/http"
)
var tmpl *template.Template

func init(){
	tmpl = template.Must(template.ParseGlob("views/home/*.html"))
}

func Splash(w http.ResponseWriter, r *http.Request) {
	// temp, err := template.ParseFiles("views/home/index.html")
	// if err != nil{
	// 	panic(err)
	// }
	// temp.Execute(w,nil)
	tmpl.ExecuteTemplate(w, "index.html", nil)
}

func Welcome(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "home.html", nil)
}
