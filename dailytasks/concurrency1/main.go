package main

import (
	"fmt"
	"sync"
	"time"
)

func Displaynum() {
	for i := 1; i < 10; i++ {
		fmt.Println("printing number:", i)
		time.Sleep(500 * time.Millisecond)
	}
}
func Displaychar() {
	for i := 'A'; i < 'J'; i++ {
		fmt.Println("printing alphabet:", string(i))
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {

	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()
		Displaynum()
	}()
	go func() {
		defer wg.Done()
		Displaychar()
	}()

	wg.Wait()
}
