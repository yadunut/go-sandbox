package main

import (
	"fmt"
	"os"
	"runtime/pprof"
	"strings"
	"sync"
)

const name = "Yanundand"

func main() {
	var counter int
	var count int
	f, err := os.Create("t.out")
	if err != nil {
		panic(err)
	}

	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()
	for {
		for {
			counter++
			v, i := loopName()
			count += i
			if strings.Contains(v, "Yadunand") {
				fmt.Println(counter)
				fmt.Println("No of routines", count)
				break
			}
		}
		counter = 0
		count = 0
	}
}

func loopName() (string, int) {
	var (
		res   []rune
		wg    sync.WaitGroup
		mx    sync.Mutex
		count int
	)
	wg.Add(len(name))
	for _, v := range name {
		go func(v rune) {
			mx.Lock()
			count++
			res = append(res, v)
			mx.Unlock()
			wg.Done()
		}(v)
	}
	wg.Wait()
	return string(res), count
}
