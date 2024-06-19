package gotemplating

import (
	"bytes"
	"io"
	"testing"

	approvals "github.com/approvals/go-approval-tests"
)

func TestRefactoredRender(t *testing.T) {
	var (
		aPost = Post{
			Title:       "Welcome to my blog",
			Description: "Introduction to my blog",
			Body:        "Welcome to my **amazing recipe blog**. I am going to write about my family recipes, and make sure I write a long, irrelevant and boring story about my family before you get to the actual instructions",
			Tags:        []string{"cooking", "family", "live-laugh-love"},
		}
	)
	renderer, err := NewPostRenderer()
	if err != nil {
		t.Fatalf("failed to get renderer %v", err)
	}
	t.Run("It converts a single post into HTML", func(t *testing.T) {
		buf := bytes.Buffer{}
		err := renderer.Render(&buf, aPost)

		if err != nil {
			t.Fatal(err)
		}

		got := buf.String()
		want := `<h1>Welcome to my blog</h1><p>Introduction to my blog</p>Tags: <ul><li>cooking</li><li>family</li><li>live-laugh-love</li></ul>`
		if got != want {
			t.Errorf("got %s but want %s", got, want)
		}
	})
}

/*
 Approval Tests - allows for easy testing of larger objects, strings and anything else that can be saved to a file
*/

func TestRender(t *testing.T) {
	var (
		aPost = Post{
			Title:       "Welcome to my blog",
			Description: "Introduction to my blog",
			Body:        "Welcome to my **amazing recipe blog**. I am going to write about my family recipes, and make sure I write a long, irrelevant and boring story about my family before you get to the actual instructions",
			Tags:        []string{"cooking", "family", "live-laugh-love"},
		}
	)
	renderer, err := NewPostRenderer()
	if err != nil {
		t.Fatalf("failed to get renderer %v", err)
	}
	t.Run("It converts a single post into HTML", func(t *testing.T) {
		buf := bytes.Buffer{}
		err := renderer.Render(&buf, aPost)

		if err != nil {
			t.Fatal(err)
		}
		approvals.VerifyString(t, buf.String())
	})
	t.Run("it renders an index if posts", func(t *testing.T) {
		buff := bytes.Buffer{}
		posts := []Post{
			{
				Title: "Hello World",
			},
			{
				Title: "Welcome Gophers",
			},
		}
		if err := renderer.RenderIndex(&buff, posts); err != nil {
			t.Fatal(err)
		}
		approvals.VerifyString(t, buff.String())
	})
}

func BenchmarkRender(b *testing.B) {
	var (
		post = Post{
			Title:       "Welcome Gophers",
			Description: "Onboarding the Gophers",
			Body:        "Creating the one forum for the gophers community",
			Tags:        []string{"go", "gowithtests", "tdd"},
		}
	)
	b.ResetTimer()
	r, err := NewPostRenderer()
	if err != nil {
		b.Fatalf("%v", err)
	}
	for i := 0; i < b.N; i++ {
		r.Render(io.Discard, post)
	}
}
