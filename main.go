package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	l := []string{
		"https://www.google.com",
		"https://stackoverflow.com",
		"https://golang.org",
		"https://github.com",
	}

	c := make(chan string)

	for _, link := range l {
		go checkLink(link, c)
	}

	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}

}

func checkLink(link string, c chan string) {

	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}
