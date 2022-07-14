package main

import (
	"fmt"
)

func main() {
	c := generator()
	receiver(c)

}

// The main function has two functions generator and receiver.
// We create a c int channel and return it from the generator function.
// The for loop running inside anonymous goroutines writing the values to the channel c.

func receiver(c <-chan int) {
	for v := range c {
		fmt.Println(v)
	}
}

func generator() <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()
	return c
}
