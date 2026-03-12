package blogrenderer

import (
	"html/template"
	"io"
)

type Post struct {
	Title, Body, Description string
	Tags                     []string
}

const postTemplate = `<h1>{{.Title}}</h1><p>{{.Description}}</p>Tags: <ul>{{range .Tags}}<li>{{.}}</li>{{end}}</ul>`

func Render(buf io.Writer, post Post) error {
	parsedTemplate, err := template.New("blog").Parse(postTemplate)
	if err != nil {
		return err
	}
	err = parsedTemplate.Execute(buf, post)
	return err
}
