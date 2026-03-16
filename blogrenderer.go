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

// Render create an HTML page from the given post variable
func (r *PostRenderer) Render(buf io.Writer, post Post) error {
	return r.templ.ExecuteTemplate(buf, "blog.html", post)
}

// PostRenderer a type that’ll hold the parsed template,
// To stop us having to re-parse the templates over and over,
type PostRenderer struct {
	templ *template.Template
}

// NewPostRenderer create a new renderer type
func NewPostRenderer() (*PostRenderer, error) {
	parsedTemplate, err := template.ParseFS(postTemplate, "templates/*.html")
	if err != nil {
		return nil, err
	}
	return &PostRenderer{templ: parsedTemplate}, nil
}
