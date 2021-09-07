//simple static file server program to find the
package main

import (
	"net/http"
	"os"
)

func main() {
	dir, _ := os.Getwd() //Getwd is similar to PWD, print working directory
	http.ListenAndServe(":3000", http.FileServer(http.Dir(dir)))
}
