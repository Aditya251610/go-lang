package main

import (
	"fmt"
	"strings"
)

func reverse(s string) string {
	n := len(s)
	var r strings.Builder
	for i := n - 1; i >= 0; i-- {
		r.WriteByte(s[i])
	}

	return r.String()
}

func main() {
	input := "Hello, World!"
	reversed := reverse(input)
	fmt.Println("Reversed string:", reversed)
	fmt.Println("Original string:", input)
}
