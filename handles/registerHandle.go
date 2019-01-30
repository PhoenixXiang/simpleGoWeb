package handles

import (
	"fmt"
	"net/http"

	"github.com/PhoenixXiang/simpleGoWeb/services"
	"github.com/PhoenixXiang/simpleGoWeb/template"
)

func registerHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		template.ExecuteTemplate(w, "register.html", View{Title: "注册"})
	} else if r.Method == "POST" {
		r.ParseForm()
		fmt.Println(r.Form)
		username := r.Form["username"][0]
		password := r.Form["password"][0]
		if username != "" && password != "" {
			u := services.QueryUser(username)
			if u == nil {
				id := services.CreateUser(username, password)
				if id > 0 {
					fmt.Println(id)
					http.Redirect(w, r, "/index", http.StatusSeeOther)
				}
			}
		}

		template.ExecuteTemplate(w, "register.html", View{Title: "注册"})
	} else {
		http.Redirect(w, r, "/error", http.StatusSeeOther)
	}

}

func init() {
	http.HandleFunc("/register", registerHandler)
}
