package view

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Main() {
	r := mux.NewRouter()
	r.HandleFunc("/", helloWorldHandler)
	http.Handle("/", r)
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello world!")
}
