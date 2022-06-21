package link

import (
	"fmt"
	"io"

	"golang.org/x/net/html"
)

// Link represents a link (<a href=".." />) in an HTML
// document
type Link struct {
	Href string
	Text string
}

// Parse will take in an HTML document and return a slice
// of parsed links
func Parse(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		panic(err)
	}

	var dfs func(*html.Node, string)
	dfs = func(n *html.Node, padding string) {
		fmt.Println(padding, n.Data)
		if n.FirstChild != nil {
			dfs(n.FirstChild, padding+"  ")
		}

		if n.NextSibling != nil {
			dfs(n.NextSibling, padding)
		}

	}
	dfs(doc, "")

	return []Link{}, nil
}
