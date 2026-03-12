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
	_, err := fmt.Fprintf(buf, `<h1>%s</h1>
<p>%s</p>
Tags: <ul><li>%s</li><li>%s</li></ul>`, post.Title, post.Description, post.Tags[0], post.Tags[1])
	return err
}
