package blogrenderer

import (
	"io"
)

type Post struct {
	Title, Body, Description string
	Tags                     []string
}

func Render(buf io.Writer, post Post) error {
	return nil
}
