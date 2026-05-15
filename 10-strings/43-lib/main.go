package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := "abc"
	s2 := strings.Clone(s1)
	fmt.Printf("s1=%s, s2=%s\n", s1, s2)

	b := strings.Builder{}
	b.Write([]byte("byte example"))
	fmt.Println(b.String())
	b.WriteString(", string example\n")
	fmt.Println(b.String())

	fmt.Println(strings.ToLower(s1))
	fmt.Println(strings.ToUpper(s1))
	fmt.Println(strings.Title(s1))

	s3 := "   " + s1 + "     "
	fmt.Println(len(s3))
	s3 = strings.TrimSpace(s3)
	fmt.Println(len(s3))

	s3 = "mario.lazzari@gmail.com"
	fmt.Println(strings.HasPrefix(s3, "mario"))
	fmt.Println(strings.HasSuffix(s3, "gmail.com"))
	fmt.Println(strings.Replace(s3, "mario", "maria", 1))

	parts := strings.Split(s3, "@")
	fmt.Printf("%+v\n", parts)
	parts = strings.Fields("Mario Lazzari")
	fmt.Printf("%+v\n", parts)

	fmt.Println(strings.Join(parts, ","))
}
