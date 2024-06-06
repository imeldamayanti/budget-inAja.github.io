package homecontroller

import (
	"KnapSack/models/makananmodel"
	"html/template"
	"net/http"
	"strconv"
	// "fmt"
	// "os"
)

var tmpl *template.Template

func init() {
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


func ShowGenerated(w http.ResponseWriter, r *http.Request) {
	frequency, _ := strconv.Atoi(r.FormValue("frequency"))
	budget, _ := strconv.ParseFloat(r.FormValue("budget"), 64)

	dt_makanan := makananmodel.GenerateData(frequency, budget)
	data := map[string]any{
		"dt_makanan": dt_makanan,
	}

	tmpl.ExecuteTemplate(w, "generatedData.html", data)
}

func Welcome(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "home.html", nil)
}
