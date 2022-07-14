package main

import "fmt"

func main() {
	ch := make(chan string)

	go func() {
		ch <- "Hello World"
		close(ch)
	}()

	for i := 0; i < 3; i++ {
		v, ok := <-ch
		fmt.Printf("Value-%q, Ok - %t\n", v, ok)
	}
}
