package v1

import (
	"errors"
	"io"
	"io/fs"
	"testing/fstest"
)

type Post struct {
	Title       string
	Description string
	Body        string
	Tags        []string
}

func NewPostsFromFS(fileSystem fstest.MapFS) ([]Post, error) {
	dir, err := fs.ReadDir(fileSystem, ".")
	if err != nil {
		return nil, err
	}
	var posts []Post
	for _, f := range dir {
		post, err := getPost(fileSystem, f.Name())
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

// Refactor it two funcs getPost and newPost and pass more sensful arguments
// func getPost(fileSystem fs.FS, f fs.DirEntry) (Post, error) {
// 	postFile, err := fileSystem.Open(f.Name())
// 	if err != nil {
// 		return Post{}, err
// 	}
// 	defer postFile.Close()

// 	postData, err := io.ReadAll(postFile)
// 	if err != nil {
// 		return Post{}, err
// 	}
// 	post := Post{Title: string(postData[7:])}
// 	return post, nil
// }

// The function `getPost` opens a file from a file system and returns a `Post` object along with any
// errors encountered.
func getPost(fileSystem fs.FS, fileName string) (Post, error) {
	postFile, err := fileSystem.Open(fileName)
	if err != nil {
		return Post{}, err
	}
	defer postFile.Close()
	return newPost(postFile)
}

// The function `newPost` reads a post file, extracts the title from the content, and returns a `Post`
// struct along with any errors encountered.
func newPost(postFile io.Reader) (Post, error) {
	postData, err := io.ReadAll(postFile)
	if err != nil {
		return Post{}, err
	}
	post := Post{Title: string(postData[7:])}
	return post, nil
}

type StubFailingFS struct {
}

func (s StubFailingFS) Open(name string) (fs.File, error) {
	return nil, errors.New("oh no i always fail")
}
