package main

import (
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

func Test_printSomething(t *testing.T) {
	stdout := os.Stdout

	r, w, err := os.Pipe()
	if err != nil {
		t.Errorf("error occured in connecting os file: %v", err)
	}

	os.Stdout = w

	var wg sync.WaitGroup
	wg.Add(1)

	go printSomething("epsilon", &wg)

	wg.Wait()

	defer w.Close()

	result, err := io.ReadAll(r)
	if err != nil {
		t.Errorf("error occured in reading file:%v", err)
	}

	output := string(result)

	os.Stdout = stdout

	if !strings.Contains(output, "epsilon") {
		t.Errorf("Expected epsilon But got %v", output)
	}

}
