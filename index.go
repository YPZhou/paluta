package paluta

import (
    "html/template"
    "net/http"
)

var t = template.Must(template.ParseFiles("templates/base.html"))

func init() {
    http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
    t.Execute(w, nil)
}