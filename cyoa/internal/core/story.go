package core

import (
	"encoding/json"
	"io"
)

func JsonStory(r io.Reader) (Story, error) {
	dec := json.NewDecoder(r)
	var story Story
	if err := dec.Decode(&story); err != nil {
		return nil, err
	}
	return story, nil
}

type Story map[string]Chapter

type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

type Option struct {
	Text    string `json:"text"`
	Chapter string `json:"arc"`
}
