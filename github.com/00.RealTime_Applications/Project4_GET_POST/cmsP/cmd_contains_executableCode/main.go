package main

import (
	"net/http"

	"github.com/00.RealTime_Applications/Project4_GET_POST/cmsP"
)

func main() {
	http.HandleFunc("/", cmsP.ServeIndex)
	http.HandleFunc("/new", cmsP.HandleNew)

	http.HandleFunc("/page/", cmsP.ServePage)
	http.HandleFunc("/post/", cmsP.ServePost)

	http.ListenAndServe(":54529", nil)
	// listener, err := net.Listen("tcp", ":0")
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("Using port:", listener.Addr().(*net.TCPAddr).Port)

	// panic(http.Serve(listener, nil))
}
