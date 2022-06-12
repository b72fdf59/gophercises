package main

import (
	"cyoa/internal/core"
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

	story, err := core.JsonStory(f)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", story)
}
