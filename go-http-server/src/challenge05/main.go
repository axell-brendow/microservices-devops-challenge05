package main

import (
	"fmt"
	"net/http"
)

func greeting(str string) string {
	return fmt.Sprintf("<b>%s</b>", str)
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, greeting("Code.education Rocks!"))
}

func main() {
	http.HandleFunc("/", greet)
	fmt.Println("Starting server on port 8000")
	http.ListenAndServe(":8000", nil)
}
