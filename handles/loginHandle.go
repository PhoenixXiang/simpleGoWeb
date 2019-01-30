package handles

import (
	"fmt"
	"net/http"

	"github.com/PhoenixXiang/simpleGoWeb/services"
	"github.com/PhoenixXiang/simpleGoWeb/template"
)

type View struct {
	Title string
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	if r.Method == "GET" {
		template.ExecuteTemplate(w, "login.html", View{Title: "登录"})
	} else if r.Method == "POST" {
		r.ParseForm()
		fmt.Println(r.Form)
		username := r.Form["username"][0]
		password := r.Form["password"][0]
		// fmt.Println(r.Form["username"][0])
		// fmt.Println(r.Form["password"][0])
		if username != "" {

			u := services.QueryUser(username)
			if u != nil && u.Password == password {
				http.Redirect(w, r, "/index", http.StatusSeeOther)
			}
		}

		template.ExecuteTemplate(w, "login.html", View{Title: "登录"})

	} else {
		http.Redirect(w, r, "/error", http.StatusSeeOther)
	}
}

func init() {
	http.HandleFunc("/login", loginHandler)
}
