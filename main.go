package main

import (
	"fmt"
	"os"
	"sync"
	"urlchecker/urlchecker"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Print("not enough arguments -- usage: main.go URL1 URL2 ...")
		return
	}
	urls := os.Args[1:]
	wg := sync.WaitGroup{}
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			urlchecker.ProcessUrl(&wg, url)
		}(url)
	}
	wg.Wait()

}
