package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	fmt.Println("Hello World")

	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Please provide a url")
		return
	}

	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	for _, url := range args {
		checkStatus(client, url)
	}
}

func checkStatus(client *http.Client, url string) {
	resp, err := client.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println(url, ": ", resp.Status)
}