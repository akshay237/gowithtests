package version2

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"strings"
	"testing/fstest"
)

const (
	titleSeprator       = "Title: "
	descriptionSeprator = "Description: "
	tagsSeprator        = "Tags: "
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
	scanner := bufio.NewScanner(postFile)

	// scanner.Scan()
	// titleLine := scanner.Text()
	// scanner.Scan()
	// descriptionLine := scanner.Text()

	// refactor above piece
	readLine := func(tagName string) string {
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), tagName)
	}

	readBody := func(bodyScanner *bufio.Scanner) string {
		bodyScanner.Scan()
		buff := bytes.Buffer{}
		for bodyScanner.Scan() {
			fmt.Fprintln(&buff, bodyScanner.Text())
		}
		return strings.TrimSuffix(buff.String(), "\n")
	}

	title := readLine(titleSeprator)
	description := readLine(descriptionSeprator)
	tags := strings.Split(readLine(tagsSeprator), ",")
	body := readBody(scanner)

	return Post{Title: title, Description: description, Tags: tags, Body: body}, nil
}

type StubFailingFS struct {
}

func (s StubFailingFS) Open(name string) (fs.File, error) {
	return nil, errors.New("oh no i always fail")
}
