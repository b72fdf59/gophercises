package link

import "io"

// Link represents a link (<a href=".." />) in an HTML
// document
type Link struct {
	Href string
	Text string
}

// Parse will take in an HTML document and return a slice
// of parsed links
func Parse(r io.Reader) ([]Link, error) {

	return []Link{}, nil
}
