package blogrenderer

import (
	"bytes"
	"embed"
	"html/template"
	"io"
	"log"
	"strings"
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

// RenderIndex ;takes an io.Writer and a slice of Post
func (r *PostRenderer) RenderIndex(buffer *bytes.Buffer, posts []Post) error {
	indexTemplate := `<ol>{{range .}}<li><a href="/post/{{.Title}}">{{.Title}}</a></li>{{end}}</ol>`

	unparsedIndexTemplate := template.New("index")

	sanitiseTitle := func(title string) string {
		log.Println(title)
		out := strings.ToLower(strings.ReplaceAll(title, " ", "-"))
		log.Println(out)
		return out
	}

	templateFuncMap := template.FuncMap{
		"sanitiseTitle": sanitiseTitle,
	}

	parsedIndexTemplate, err := unparsedIndexTemplate.Funcs(templateFuncMap).Parse(indexTemplate)
	if err != nil {
		return err
	}

	return parsedIndexTemplate.Execute(buffer, posts)
}

// NewPostRenderer create a new renderer type
func NewPostRenderer() (*PostRenderer, error) {
	parsedTemplate, err := template.ParseFS(postTemplate, "templates/*.html")
	if err != nil {
		return nil, err
	}
	return &PostRenderer{templ: parsedTemplate}, nil
}
