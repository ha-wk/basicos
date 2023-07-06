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
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	for l := range c {
		go func(link string) { //link is the value that we want this func to receive
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l) //to say that this l is used inside function
	}

	/* for i:=0;i<len(links);i++{
		fmt.Println(<-c)
	} */
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

/*package main

import "net/http"
import "fmt"
import "sync"

func main() {

	var wg sync.WaitGroup

	links := []string{
		"http://google.com",
		"http://stackoverflow.com",
		"http://facebook.com",
		"http://golang.org",
	}
	//c:=make(chan string)

	wg.Add(len(links))

	for _, link := range links {

		defer wg.Done()
		checklink(link)

	}
	wg.Wait()
	fmt.Println("All Tested Succesfully")
}

func checklink(link string) {

	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "is down currently")
		return
	}
	fmt.Println(link, "is abs working fine")
}

*/
