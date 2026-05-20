package main

import (
	"encoding/base64"
	"fmt"
	"log"
)

func main() {
	data := "Welcome to Go!"
	encoded := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(encoded)

	endodedStr := "V2VsY29tZSB0byBHbyE="
	decodedStr, err := base64.StdEncoding.DecodeString(endodedStr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(decodedStr))

	rawData := []byte{0xAA, 0xDE}
	rawStr := base64.StdEncoding.EncodeToString(rawData)
	fmt.Println(rawStr)

	b64str := "qt4="
	decodedStr, err = base64.StdEncoding.DecodeString(b64str)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(decodedStr))
}
