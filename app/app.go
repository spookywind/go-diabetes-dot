package app

import (
	"fmt"
	"net/http"
)

var MY_HTML string = `
<!DOCTYPE html>
<html>
<head>
	<title>This is Hello HTML page</title>
</head>
<body>
 	<h1>Hello HTML!</h1>
</body>
</html>
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
