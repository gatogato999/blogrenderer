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

	postRenderer, err := blogrenderer.NewPostRenderer()
	if err != nil {
		t.Fatal(err)
	}

	t.Run("it converts a single post into HTML", func(t *testing.T) {
		buf := bytes.Buffer{}
		err = postRenderer.Render(&buf, aPost)
		if err != nil {
			t.Fatal(err)
		}
		approvals.VerifyString(t, buf.String())
	})

	t.Run("it renders an index of posts", func(t *testing.T) {
		buf := bytes.Buffer{}
		posts := []blogrenderer.Post{{Title: "Hello World"}, {Title: "Hello World 2"}}
		if err := postRenderer.RenderIndex(&buf, posts); err != nil {
			t.Fatal(err)
		}
		got := buf.String()
		want := `<ol><li><a href="/post/hello-world">Hello World</a></li><li><a href="/post/hello-world-2">Hello World 2</a></li></ol>`
		if got != want {
			t.Errorf("\ngot %q \nwant %q", got, want)
		}
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
	postRend, err := blogrenderer.NewPostRenderer()
	if err != nil {
		b.Fatal(err)
	}
	for b.Loop() {
		_ = postRend.Render(io.Discard, aPost)
	}
}
