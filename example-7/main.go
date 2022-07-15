package main

import "fmt"

func main() {
	chanFirth := make(chan string)
	chanSecond := make(chan string)
	go sourceGopher(chanFirth)
	go removeDuplicates(chanFirth, chanSecond)
	printGopher(chanSecond)
}

func sourceGopher(down chan string) {
	for _, v := range []string{"a", "b", "b", "c", "d", "d", "d", "e", "r"} {
		down <- v
	}
	close(down)
}

func removeDuplicates(up, down chan string) {
	prev := ""
	for v := range up {
		if v != prev {
			down <- v
			prev = v
		}
	}
	close(down)
}
func printGopher(up chan string) {
	for v := range up {
		fmt.Println(v)
	}
}
