package main

import (
	"fmt"
	"time"
)

func printSomething(s string) {
	fmt.Println(s)
}

func main() {
	go printSomething("this is first call!")

	time.Sleep(1 * time.Millisecond)

	printSomething("this is second call!")
}
