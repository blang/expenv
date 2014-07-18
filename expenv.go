package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: expenv inputfile")
		return
	}

	b, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	outStr := os.ExpandEnv(string(b))
	ioutil.WriteFile(os.Args[1], []byte(outStr), 0)
}
