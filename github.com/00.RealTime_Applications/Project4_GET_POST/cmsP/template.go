//template - to Hold HTML template
package cmsP

import (
	"html/template"
	"time"
)

//Tmpl is an exported variable
var Tmpl = template.Must(template.ParseGlob("../templates/*"))

type Page struct {
	Title   string
	Content string
	Posts   []*Post
}
type Post struct {
	Title         string
	Content       string
	DatePublished time.Time
	Comments      []*Comment
}

type Comment struct {
	Author        string
	Comment       string
	DatePublished time.Time
}
