package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type TestCase struct {
	Name    string                 `json:"name"`
	Method  string                 `json:"method"`
	URL     string                 `json:"url"`
	Headers map[string]string      `json:"headers"`
	Body    map[string]interface{} `json:"body,omitempty"`
}

func loadTestCases(filePath string) ([]TestCase, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var testCases []TestCase
	if err := json.Unmarshal(data, &testCases); err != nil {
		return nil, err
	}

	return testCases, nil
}

func sendRequest(tc TestCase) error {
	var req *http.Request
	var err error

	var bodyReader io.Reader
	if tc.Method == "POST" || tc.Method == "PUT" || tc.Method == "PATCH" {
		bodyBytes, err := json.Marshal(tc.Body)
		if err != nil {
			return err
		}
		bodyReader = bytes.NewBuffer(bodyBytes)
	}

	req, err = http.NewRequest(tc.Method, tc.URL, bodyReader)
	if err != nil {
		return err
	}

	// Add headers
	for k, v := range tc.Headers {
		req.Header.Set(k, v)
	}

	client := &http.Client{Timeout: 10 * time.Second}
	start := time.Now()
	resp, err := client.Do(req)
	duration := time.Since(start)

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	fmt.Printf("✅ [%s] %s %s\n", tc.Method, tc.Name, tc.URL)
	fmt.Printf("   → Status: %d, Time: %v\n", resp.StatusCode, duration)

	return nil
}

func main() {
	testCases, err := loadTestCases("/home/aditya/Documents/Web Development/go-lang/mini project 1/test.json")
	if err != nil {
		fmt.Println("Error loading test cases:", err)
		return
	}

	for _, tc := range testCases {
		if err := sendRequest(tc); err != nil {
			fmt.Printf("❌ Error in '%s': %v\n", tc.Name, err)
		}
	}
}
