package main

import (
	"cyoa/internal/core"
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

func main() {
	filename := flag.String("file", "gopher.json", "The JSON file with the choose your own adventure story")
	flag.Parse()
	fmt.Printf("Using the file %s\n", *filename)

	f, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}

	dec := json.NewDecoder(f)
	var story core.Story
	if err := dec.Decode(&story); err != nil {
		panic(err)
	}
	fmt.Printf("%+v", story)
}
