package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Time{}
	fmt.Println(t.Sub(time.Time{}) == 0)
}
