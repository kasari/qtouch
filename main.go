package main

import (
	"fmt"
	"os"
	"path"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("qtouch path/to/filename")
		os.Exit(1)
	}

	filepath := os.Args[1]
	dir, _ := path.Split(filepath)

	os.MkdirAll(dir, 0755)

	f, err := os.Create(filepath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	f.Close()
}
