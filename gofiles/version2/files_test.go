package version2

import (
	"testing"
	"testing/fstest"

	"github.com/google/go-cmp/cmp"
)

func TestNewBlogPosts(t *testing.T) {
	const (
		firstBody = `Title: Hello, TDD world!
Description: First post on our wonderful blog
Tags: tdd,go
---
Hello world!`
		secondBody = `Title: Hello, Go Developers!
Description: Second post on our wonderful blog
Tags: tdd,go,gotests
---
Hello Developers!`
	)

	fs := fstest.MapFS{
		"helloworld1.md": {Data: []byte(firstBody)},
		"helloworld2.md": {Data: []byte(secondBody)},
	}

	posts, err := NewPostsFromFS(fs)
	if err != nil {
		t.Fatal(err)
	}

	// first post test
	got1 := posts[0]
	want1 := Post{
		Title:       "Hello, TDD world!",
		Description: "First post on our wonderful blog",
		Tags:        []string{"tdd", "go"},
		Body:        "Hello world!",
	}
	assertPost(t, got1, want1)

	// second post test
	got2 := posts[1]
	want2 := Post{
		Title:       "Hello, Go Developers!",
		Description: "Second post on our wonderful blog",
		Tags:        []string{"tdd", "go", "gotests"},
		Body:        "Hello Developers!",
	}
	assertPost(t, got2, want2)

	if len(posts) != len(fs) {
		t.Errorf("got %d posts, wanted %d posts", len(posts), len(fs))
	}
}

func assertPost(t *testing.T, got, want Post) {
	t.Helper()
	if !cmp.Equal(got, want) {
		t.Errorf("got:%v, wanted:%v", got, want)
	}
}
