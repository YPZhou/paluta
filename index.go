package paluta

import (
    "html/template"
    "net/http"
)

var tHome = template.Must(template.ParseFiles("templates/base.html", "templates/home.html"))
var tGalery = template.Must(template.ParseFiles("templates/base.html", "templates/galery.html"))
var tProgress = template.Must(template.ParseFiles("templates/base.html", "templates/progress.html"))
var tJoin = template.Must(template.ParseFiles("templates/base.html", "templates/join.html"))

func init() {
    http.HandleFunc("/", handler)
    http.HandleFunc("/home", handler)
    http.HandleFunc("/galery", handler)
    http.HandleFunc("/progress", handler)
    http.HandleFunc("/join", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	strPath := r.URL.Path[1:]
	switch strPath {
		case "home" :
			tHome.Execute(w, nil)
		case "galery" :
			tGalery.Execute(w, nil)
		case "progress" :
			tProgress.Execute(w, nil)
		case "join" :
			tJoin.Execute(w, nil)
		default :
			tHome.Execute(w, nil)
	}
}