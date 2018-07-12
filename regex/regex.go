package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

var doneRegex = regexp.MustCompile(`\[([xX\s])\]`)

//var priorityRegex = regexp.MustCompile(``)
//var dateRegex = regexp.MustCompile(``)

func main() {
	f, err := os.Open("sample.txt")
	scn := bufio.NewScanner(f)

	for scn.Scan() {
		s := doneRegex.FindStringSubmatch(scn.Text())
		if s != nil {
			fmt.Println(s)
		} else {
			fmt.Println("err")
		}
	}
	if scn.Err() != nil {
		panic(err)
	}

}
