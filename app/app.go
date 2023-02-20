package app

import (
	"fmt"
	"net/http"
)

var MY_HTML string = `
HELLO WORLD
`

func Run() {
	http.HandleFunc("/go-diabetes-dot", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Println("GET")
	}
	fmt.Fprintf(w, "%s\n", MY_HTML)
}
