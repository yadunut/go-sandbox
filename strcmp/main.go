package main

import "fmt"

func main() {
	i1 := "string1"
	i2 := "string2"
	fmt.Println(Strcmp(i1, i2))
}

func Strcmp(i1, i2 string) int {
	var r int
	for i := range i1 {
		if i1[i] != i2[i] {
			r++
		}
	}

	return r
}
