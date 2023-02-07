package main

import (
	"fmt"
	"sync"
)

func printSomething(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(s)
}

func main() {
	var wg sync.WaitGroup

	str := []string{
		"alpha",
		"beta",
		"gamma",
		"theta",
		"omega",
	}

	wg.Add(len(str))

	for i, v := range str {
		go printSomething(fmt.Sprintf("index %d, value %s", i, v), &wg)
	}

	wg.Wait()
	wg.Add(1)

	printSomething("this is last line", &wg)
}
