package main

import (
	"log"
	"net/http"
	"KnapSack/config"
	"KnapSack/controllers/aboutcontroller"
	"KnapSack/controllers/homecontroller"
	"KnapSack/controllers/makanancontroller"
	"KnapSack/models/makananmodel"
)

// var tmpl *template.Template

// func init(){
// 	tmpl = template.Must(template.ParseGlob("views/home/*.html"))
// }

// func homeHandler(w http.ResponseWriter, r *http.Request) {
// 	tmpl.ExecuteTemplate(w, "index.html", nil)
// }

func main() {
	config.ConnectDB()
	makananmodel.GetAll()

	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets", fs))
	// http.HandleFunc("/", homeHandler)

	http.HandleFunc("/", homecontroller.Splash)
	http.HandleFunc("/generate", homecontroller.Generate)
	http.HandleFunc("/generated-data", homecontroller.ShowGenerated)
	http.HandleFunc("/data", makanancontroller.Index)
	http.HandleFunc("/about", aboutcontroller.Index)


	log.Println("Server running on port 3001")
	http.ListenAndServe("127.0.0.1:3001", nil)
}

// go mod init
// go get -u "github.com/go-sql-driver/mysql"
// go run .
// minta tolong buat hubungin depan -> algoritmanya