package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Event struct {
	Type string `json:"type"`
	Repo struct {
		Name string `json:"name"`
	} `json:"repo"`
	Payload struct {
		Commits []struct {
			Message string `json:"message"`
		} `json:"commits"`
	} `json:"payload"`
}

func main() {
	if len(os.Args) < 1 {
		fmt.Println("Username required")
		return
	}
	username := os.Args[1]
	url := fmt.Sprintf("https://api.github.com/users/%s/events", username)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching user events:", err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Failed to fetch events. HTTP Status: %d\n", resp.StatusCode)
		return
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}
	var events []Event
	if err := json.Unmarshal(body, &events); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
	fmt.Printf("Recent events for user '%s':\n", username)
	for _, event := range events {
		describeEvent(event)
	}
}
