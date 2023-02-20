package app

import (
	"fmt"
	"go-diabetes-dot/html"
	"net/http"
)

func Run() {
	http.HandleFunc("/go-diabetes-dot", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Println("GET")
	}
	fmt.Fprintf(w, "%s\n", html.Html())
}
