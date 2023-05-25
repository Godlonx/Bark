package bark

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func WebServer() {
	http.HandleFunc("/", Home)
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets", fs))

	fmt.Println("Starting server at port 8871 : http://localhost:8871")
	err := http.ListenAndServe(":8871", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func Home(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("template/Home.html"))
	if r.Method == http.MethodPost {
		test := r.FormValue("test")
		tagInfo := strings.Split(test, " ")
		for _, val := range tagInfo {
			if val[0] == '#' {
				print(val)
			}
		}
	}
	err := tmpl.Execute(w, nil)
	if err != nil {
		log.Fatal(err)
	}
}
