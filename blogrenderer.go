package blogrenderer

import (
	"fmt"
	"io"
)

type Post struct {
	Title, Body, Description string
	Tags                     []string
}

func Render(buf io.Writer, post Post) error {
	_, err := fmt.Fprintf(buf, `<h1>%s</h1>`, post.Title)
	return err
}
