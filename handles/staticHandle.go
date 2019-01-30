package handles

import "net/http"

func Static (path, dir string)  {
	http.Handle(path, http.StripPrefix(path, http.FileServer(http.Dir(dir))))
}
