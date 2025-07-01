package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email,omitempty"`
}

func main() {
	person := Person{
		Name:  "Aditya",
		Age:   22,
		Email: "aditya.dev2516@gmail.com",
	}

	jsonData, err := json.MarshalIndent((person), "", "  ")
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	err = os.WriteFile("person.json", jsonData, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("Struct Successfully writen to person.json")

}
