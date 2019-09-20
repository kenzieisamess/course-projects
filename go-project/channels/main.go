package main

import (
	"fmt"
	"net/http"
)

func checkLink(link string) {
	_, e := http.Get(link)

	if e != nil {
		fmt.Println(link, "might be down!")
		return
	}

	fmt.Println(link, "is up!")
}

func main() {
	links := []string{
		"http://google.com",
		"http://stackoverflow.com",
	}

	for _, l := range links {
		go checkLink(l)
	}
}
