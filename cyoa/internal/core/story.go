package core

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"text/template"
)

func init() {
	data, err := os.ReadFile("static/cyoa.html")
	if err != nil {
		panic(err)
	}
	tpl = template.Must(template.New("Story").Parse(string(data)))
}

var tpl *template.Template

func NewHandler(s Story) http.Handler {
	return handler{s: s}
}

type handler struct {
	s Story
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := tpl.Execute(w, h.s["intro"])
	if err != nil {
		panic(err)
	}
}

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
