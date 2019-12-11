package main

import (
	"fmt"
	"io"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, `Hello (Real) World`)
}
func main() {
	fmt.Println("Serving on 8000 port in environment %s:::::", EnvSpecificValue)
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8000", nil)

}
