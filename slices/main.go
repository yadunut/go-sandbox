package main

import "fmt"

type arr struct {
	val1 string
	val2 string
}

type arrs []arr

func main() {
	as := arrs{
		{"val1", "val2"},
		{"val3", "val4"},
		{"val5", "val6"},
	}
	fmt.Println(as)
	bs := arrs{
		{"val1", "val2"},
		{"val3", "val4"},
		{"val5", "val6"},
	}

	if as == bs {
		fmt.Println("Equal")
	}
}
