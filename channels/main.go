package main

import (
	"fmt"
	"net/http"
)

func checkLink(link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		return
	}
	fmt.Println(link, "is up!")
}

func main() {
	links := []string{
		"https://google.com",
		"https://amazon.com",
		"https://twitter.com",
		"https://golang.org",
	}
	for _, link := range links {
		checkLink(link)
	}
}
