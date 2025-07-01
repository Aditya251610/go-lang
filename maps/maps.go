package main

import "fmt"

func main() {
	countries := map[string]string{
		"IN":  "India",
		"US":  "United States",
		"UK":  "United Kingdom",
		"PAK": "Pakistan",
	}

	for key, value := range countries {
		fmt.Println("Country Code:", key, "-> Country Name:", value)
	}
}
