package main

import (
	"fmt"
	"os"
)

const (
	dir  = "dir"
	file = "file.txt"
)

func main() {
	check(dir)
	check(file)
}

func check(file string) {
	f, err := os.Stat(file)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("%s Is Directory: %t Is File: %t\n", file, f.Mode().IsDir(), f.Mode().IsRegular())

}
