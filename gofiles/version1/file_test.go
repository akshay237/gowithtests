package v1

import (
	"testing"
	"testing/fstest"

	"github.com/google/go-cmp/cmp"
)

func TestNewBlogPosts(t *testing.T) {
	fs := fstest.MapFS{
		"helloworld.md": {Data: []byte("Title: Post 1")},
	}
	posts, err := NewPostsFromFS(fs)
	got := posts[0]
	want := Post{Title: "Post 1"}
	assertPost(t, got, want)
	if err != nil {
		t.Fatal(err)
	}
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
