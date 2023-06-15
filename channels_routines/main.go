package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func findurlsize(url string, nums chan int) {
	defer wg.Done()

	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)

	}

	//defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	nums <- len(body)


func main() {
	wg.Add(1)

	nums := make(chan int)
	url := "https://www.golangprograms.com/goroutines.html"
	go findurlsize(url, nums)
	ans := <-nums
	fmt.Println(ans)
	wg.Wait()
	close(nums)
}
