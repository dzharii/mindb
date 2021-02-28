package webapi

import (
	"fmt"
	"net/http"
)

func getIndex(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-type", "text/html")
	html := `
	<html>
	<body style="background-color: pink;">
		<h1>hello Vorld!</h1>
	<body>
	<html>

`
	fmt.Fprintf(w, html)
}

// StartServer starts the server
func StartServer(addr string) error {
	http.HandleFunc("/", getIndex)
	err := http.ListenAndServe(addr, nil)
	return err
}
