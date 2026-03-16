package blogrenderer

import (
	"embed"
	"html/template"
	"io"
)

// Post represent a social media post
type Post struct {
	Title, Body, Description string
	Tags                     []string
}

//go:embed "templates/*"
var postTemplate embed.FS

// Render create a HTML page from the given post variable
func Render(buf io.Writer, post Post) error {
	parsedTemplate, err := template.ParseFS(postTemplate, "templates/*.html")
	if err != nil {
		return err
	}
	err = parsedTemplate.Execute(buf, post)
	return err
}
