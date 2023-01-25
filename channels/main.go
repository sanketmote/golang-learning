package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{"https://www.druva.com/", "https://www.amazon.org", "https://www.sanket.com", "https://www.google.com/", "https://www.amazon.com/"}
	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	for l := range c {
		go func(l string) {
			time.Sleep(5 * time.Second)
			go checkLink(l, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	// time.Sleep(5 * time.Second)
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "Might be down.")
		c <- link
		return
	}
	fmt.Println(link, "is up.")

	c <- link
}

// package main

// import "fmt"

// func main() {
// 	c := make(chan string)
// 	for i := 0; i < 4; i++ {
// 		go printString("Hello there!", c)
// 	}

// 	for s := range c {
// 		fmt.Println(s)
// 	}
// }

// func printString(s string, c chan string) {
// 	fmt.Println(s)
// 	c <- "Done printing."
// }

// package main

// import "fmt"

// func main() {
//  c := make(chan string)

//  for i := 0; i < 4; i++ {
//      go printString("Hello there!", c)
//  }

//  for {
//      fmt.Println(<- c)
//  }
// }

// func printString(s string, c chan string) {
//  fmt.Println(s)
//  c <- "Done printing."
// }
