package main

import "fmt"

func main() {
	ch := make(chan int)
	ch <- 3
	fmt.Println(<-ch)
}
