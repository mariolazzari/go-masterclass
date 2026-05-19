package main

import (
	"log"
	"os"
)

func main() {
	tmpFile, err := os.CreateTemp("", "log.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		log.Println("Removing temp files...", tmpFile.Name())
		_ = os.Remove(tmpFile.Name())
	}()

	_, err = tmpFile.WriteString("Ciao Mario")
	if err != nil {
		log.Fatal(err)
	}

	tmpDir, err := os.MkdirTemp("logs", "log.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		log.Println("Removing temp files...", tmpDir)
		_ = os.Remove(tmpDir)
	}()
}
