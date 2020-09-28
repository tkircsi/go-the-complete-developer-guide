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
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
		"http://fundamenta.hu",
		"https://ebank.fundamenta.hu",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	// for i := 0; i < len(links); i++ {
	// 	<-c
	// }

	// for {
	// 	go checkLink(<-c, c)
	// }

	for l := range c {
		// Function literal (Lambda function)
		go func(l string) {
			time.Sleep(3 * time.Second)
			checkLink(l, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	start := time.Now()
	resp, err := http.Get(link)
	elapsed := time.Since(start).Milliseconds()

	if err != nil {
		fmt.Println(link, "might be down")
		c <- link
		// c <- link + " might be down"
		return
	}
	fmt.Println(link, "--", resp.Status, "--", elapsed, "ms")
	c <- link
	// c <- link + " is OK and respond " + resp.Status
	// c <- fmt.Sprintf("%s is OK and respond %s", link, resp.Status)
}
