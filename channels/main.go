package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://uwaterloo.ca",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string) //creat a new channel

	for _, link := range links {
		go checkLink(link, c) //go routine / thread
	}
	// // fmt.Println(<-c) //blocking code, white for the response from the channel
	// for { //infinte loop
	// 	//fmt.Println(<-c)
	// 	go checkLink(<-c, c)
	// }
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
