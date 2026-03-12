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
	_, err := fmt.Fprintf(buf, `<h1>%s</h1> <p>%s</p>`, post.Title, post.Description)
	if err != nil {
		return err
	}

	_, err = fmt.Fprint(buf, `Tags: <ul>`)
	if err != nil {
		return err
	}

	for _, tag := range post.Tags {
		_, err := fmt.Fprintf(buf, `<li>%s</li>`, tag)
		if err != nil {
			return err
		}
	}

	_, err = fmt.Fprint(buf, `</ul>`)
	if err != nil {
		return err
	}
	return nil
}
