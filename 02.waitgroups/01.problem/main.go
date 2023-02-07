package main

import (
	"fmt"
	"time"
)

func printSomething(s string) {
	fmt.Println(s)
}

func main() {
	str := []string{
		"alpha",
		"beta",
		"gamma",
		"theta",
		"omega",
	}

	for i, v := range str {
		go printSomething(fmt.Sprintf("index %d, value %s", i, v))
	}

	time.Sleep(1 * time.Millisecond)
}
