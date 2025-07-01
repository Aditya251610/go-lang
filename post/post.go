package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	url := "https://jsonplaceholder.typicode.com/posts"

	data := map[string]interface{}{
		"title":  "foo",
		"body":   "bar",
		"userId": 1,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("Error marshaling JSON: %v\n", err)
		return
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("Error creating request: %v\n", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error performing POST request: %v\n", err)
	}

	defer resp.Body.Close()
	fmt.Printf("Status Code: %d\n", resp.StatusCode)
}
