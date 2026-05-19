package main

import (
	"log"
	"os"
	"path/filepath"
)

func main() {
	dir := "download"

	// mkdir
	if err := os.Mkdir(dir, 0755); err != nil {
		log.Fatal(err)
	}
	// del
	if err := os.Remove(dir); err != nil {
		log.Fatal(err)
	}

	dir = "download/static/images"

	// mkdir -p
	if err := os.MkdirAll(filepath.Clean(dir), 0755); err != nil {
		log.Fatal(err)
	}

	// del
	dir = "download"
	if err := os.RemoveAll(filepath.Clean(dir)); err != nil {
		log.Fatal(err)
	}
}
