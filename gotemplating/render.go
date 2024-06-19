package gotemplating

import (
	"embed"
	"fmt"
	"io"
	"text/template"

	"github.com/gomarkdown/markdown/parser"
)

/*
 text/template and html/template are two templating package provided by go. They both share the same interface.
 html - implements data driven templates for generating HTML output safe against code injection.
*/

var (
	//go:embed "templates/*"
	postTemplate embed.FS
)

type PostRenderer struct {
	template *template.Template
	mdParser *parser.Parser
}

func NewPostRenderer() (*PostRenderer, error) {
	template, err := template.ParseFS(postTemplate, "templates/*.gohtml")
	if err != nil {
		return nil, err
	}
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs
	parser := parser.NewWithExtensions(extensions)
	return &PostRenderer{
		template: template,
		mdParser: parser,
	}, nil
}

func RefactoredRender(w io.Writer, p Post) error {
	_, err := fmt.Fprintf(w, "<h1>%s</h1><p>%s</p>", p.Title, p.Description)
	if err != nil {
		return err
	}

	_, err = fmt.Fprint(w, "Tags: <ul>")
	if err != nil {
		return err
	}

	for _, tag := range p.Tags {
		_, err := fmt.Fprintf(w, "<li>%s</li>", tag)
		if err != nil {
			return err
		}
	}

	_, err = fmt.Fprint(w, "</ul>")
	if err != nil {
		return err
	}
	return nil
}

func (r *PostRenderer) Render(w io.Writer, p Post) error {
	return r.template.ExecuteTemplate(w, "blog.gohtml", p)
}

func (r *PostRenderer) RenderIndex(w io.Writer, posts []Post) error {
	return r.template.ExecuteTemplate(w, "index.gohtml", posts)
}
