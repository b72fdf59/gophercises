package core

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
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
	path := strings.TrimSpace(r.URL.Path)
	if path == "" || path == "/" {
		path = "/intro"
	}
	path = path[1:]

	if chapter, ok := h.s[path]; ok {
		err := tpl.Execute(w, chapter)
		if err != nil {
			log.Printf("%v", err)
			http.Error(w, "Something went wrong...", http.StatusInternalServerError)
		}
		return
	}
	http.Error(w, "Chapter not found", http.StatusNotFound)
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
