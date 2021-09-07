package main

import (
	"flag"
	"net/http"
	"os"
)

func main() {
	var dirr string

	port := flag.String("port", "3000", "port to serve HTTP on")
	path := flag.String("path", "", "path to serve")
	flag.Parse()

	if *path == "" {
		dirr, _ = os.Getwd()
	} else {
		dirr = *path
	}

	http.ListenAndServe(":"+*port, http.FileServer(http.Dir(dirr)))
}

//execute in the terminal : go run fs.go -port=5158 -path=site
