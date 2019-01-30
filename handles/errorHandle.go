package handles

import (
	"fmt"
	"net/http"

	"github.com/PhoenixXiang/simpleGoWeb/template"
)

type ErrorView struct {
	Title   string
	Code    int
	Message string
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL)
	http.Redirect(w, r, "/error", http.StatusSeeOther)
}

func errorHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Println(r.URL)
	template.ExecuteTemplate(w, "error.html", ErrorView{Title: "错误", Code: http.StatusNotFound, Message: "The page you’re looking for was not found."})
}

func init() {
	http.HandleFunc("/", notFoundHandler)
	http.HandleFunc("/error", errorHandler)
}
