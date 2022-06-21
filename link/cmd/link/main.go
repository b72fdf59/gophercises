package main

import (
	"flag"
	"fmt"
	"link/internal/link-core"
	"os"
)

func main() {
	filename := flag.String("file", "examples/ex1.html", "The HTML file to parse")
	flag.Parse()
	fmt.Printf("Using the file %s\n", *filename)

	f, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}

	links, err := link.Parse(f)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", links)
}
