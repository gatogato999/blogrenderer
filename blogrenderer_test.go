package blogrenderer_test

import (
	"bytes"
	"io"
	"testing"

	approvals "github.com/approvals/go-approval-tests"

	"github.com/gatogato999/blogrenderer"
)

func TestRender(t *testing.T) {
	aPost := blogrenderer.Post{
		Title:       "hello world",
		Body:        "This is a post",
		Description: "This is a description",
		Tags:        []string{"go", "tdd"},
	}
	t.Run("it converts a single post into HTML", func(t *testing.T) {
		buf := bytes.Buffer{}
		err := blogrenderer.Render(&buf, aPost)
		if err != nil {
			t.Fatal(err)
		}
		approvals.VerifyString(t, "Hello World!\n")
	})
}

// To see the impact of not doing this parsing over and over,
// we can use the benchmarking tool to see how fast our function is.
// run it with : go test -bench=.
func BenchmarkRender(b *testing.B) {
	aPost := blogrenderer.Post{
		Title:       "hello world",
		Body:        "This is a post",
		Description: "This is a description",
		Tags:        []string{"go", "tdd"},
	}
	for b.Loop() {
		_ = blogrenderer.Render(io.Discard, aPost)
	}
}
