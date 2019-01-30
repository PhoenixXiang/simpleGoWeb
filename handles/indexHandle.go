package handles

import (
	"fmt"
	"net/http"

	"github.com/PhoenixXiang/simpleGoWeb/template"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	if r.Method == "GET" {
		template.ExecuteTemplate(w, "index.html", View{Title: "主页"})
	} else if r.Method == "POST" {
		r.ParseForm()
		fmt.Println(r.Form)
		fmt.Println(r.Form["username"][0])
		fmt.Println(r.Form["password"][0])
	}
}

func init() {
	http.HandleFunc("/index", indexHandler)
}
