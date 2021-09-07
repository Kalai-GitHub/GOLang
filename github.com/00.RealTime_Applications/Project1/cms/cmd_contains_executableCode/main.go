package main

import (
	"os"

	"github.com/00.RealTime_Applications/Project1/cms"
)

func main() {
	p := &cms.Page{
		Title:   "Hello, world!",
		Content: "This is the body of our webpage",
	}
	cms.Tmpl.ExecuteTemplate(os.Stdout, "index", p) //reading HTTP response to OS
}
