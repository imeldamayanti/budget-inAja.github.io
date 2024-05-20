package main

import (
	"log"
	"net/http"
	"tubes_sa/controllers/homecontroller"
)

// var tmpl *template.Template

// func init(){
// 	tmpl = template.Must(template.ParseGlob("views/home/*.html"))
// }

// func homeHandler(w http.ResponseWriter, r *http.Request) {
// 	tmpl.ExecuteTemplate(w, "index.html", nil)
// }

func main() {
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets", fs))
	// http.HandleFunc("/", homeHandler)

	http.HandleFunc("/", homecontroller.Splash)
	// http.HandleFunc("/home", homecontroller.Welcome)

	log.Println("Server running on port 3001")
	http.ListenAndServe("127.0.0.1:3001", nil)
}