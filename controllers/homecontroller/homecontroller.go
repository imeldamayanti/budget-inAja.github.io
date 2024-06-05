package homecontroller

import (
	"html/template"
	"net/http"
	"KnapSack/models/makananmodel"
	"strconv"
	"fmt"
	// "os"
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

// Coba aja kalau kamu bisa Redirect :3
// func Generate(w http.ResponseWriter, r *http.Request) {
// 	frequency, _ := strconv.Atoi(r.FormValue("frequency"))
// 	budget, _ := strconv.Atoi(r.FormValue("budget"))

// 	newURL := fmt.Sprintf("/generated-data?frequency=%d&budget=%d", frequency, budget)

// 	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3001" + newURL)
// 	http.Redirect(w, r, "http://localhost:3001" + newURL, http.StatusSeeOther)

// 	// w.WriteHeader(http.StatusFound)
// 	// w.Header().Set("Location", newURL)
// 	// return
// }

func ShowGenerated(w http.ResponseWriter, r *http.Request) {
	frequency, _ := strconv.Atoi(r.FormValue("frequency"))
	budget, _ := strconv.ParseFloat(r.FormValue("budget"), 64)

	dt_makanan := makananmodel.GenerateData(frequency, budget)
	data := map[string]any{
		"dt_makanan": dt_makanan,
	}

	tmpl.ExecuteTemplate(w, "data.html", data)
}

func Welcome(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "home.html", nil)
}
