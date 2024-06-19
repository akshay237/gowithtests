package gotemplating

import (
	"html/template"
	"strings"

	"github.com/gomarkdown/markdown"
)

type Post struct {
	Title, sanitiseTitle, Description, Body string
	Tags                                    []string
}

func (p Post) SanitisedTitle() string {
	return strings.ToLower(strings.Replace(p.Title, " ", "-", -1))
}

type postViewModel struct {
	Post
	HTMLBody template.HTML
}

func newPostVM(p Post, r *PostRenderer) postViewModel {
	htmlBody := template.HTML(markdown.ToHTML([]byte(p.Body), r.mdParser, nil))
	return postViewModel{
		Post:     p,
		HTMLBody: htmlBody,
	}
}
