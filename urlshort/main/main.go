package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"urlshort"
)

var pathsToUrls = map[string]string{
	"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
	"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
}

func main() {
	yamlFileName := flag.String("yamlFileName", "path.yaml", "This is the file path for the yaml file to load data from")
	flag.Parse()

	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	// Build the YAMLHandler using the mapHandler as the
	// fallback
	yaml, err := loadFileContents(*yamlFileName)
	if err != nil {
		panic(err)
	}
	yamlHandler, err := urlshort.YAMLHandler([]byte(yaml), mapHandler)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}

func loadFileContents(fileName string) (string, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
