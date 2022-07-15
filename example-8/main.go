package main

import (
	"fmt"
	"strings"
)

func main() {
	chanFirth := make(chan string)
	chanSecond := make(chan string)
	go sourceGopher(chanFirth)
	go splitWords(chanFirth, chanSecond)
	printGopher(chanSecond)
}

func sourceGopher(down chan string) {
	for _, v := range []string{"hello world!", "a bad apple", "goodbue all"} {
		down <- v
	}
	close(down)
}

func splitWords(up, down chan string) {
	for v := range up {
		for _, word := range strings.Fields(v) {
			down <- word
		}
	}
	close(down)
}
func printGopher(up chan string) {
	for v := range up {
		fmt.Println(v)
	}
}
