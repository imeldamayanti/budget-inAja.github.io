package makanancontroller

import (
	"net/http"
	"text/template"
	"KnapSack/models/makananmodel"
)

func Index(w http.ResponseWriter, r *http.Request) {
	dt_makanan := makananmodel.GetAll()
	data := map[string]any{
		"dt_makanan": dt_makanan,
	}

	temp, err := template.ParseFiles("views/home/data.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)
}
