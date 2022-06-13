package main

import (
	"cyoa/internal/core"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := flag.Int("port", 3000, "the port to sart the CYOA web application on")
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

	h := core.NewHandler(story, nil)
	fmt.Printf("Starting the server on the port: %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))

}
