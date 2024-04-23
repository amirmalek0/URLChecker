package main

import (
	"fmt"
	"net/http"
	"os"
	"regexp"
	"sync"
)

func checkUrlValidity(url string) (bool, error) {
	pattern, err := regexp.Compile("^https?://[^/]+")
	if err != nil {
		return false, err
	}
	return pattern.MatchString(url), nil
}
func checkUrlAvailability(url string) (bool, string) {
	response, err := http.Get(url)
	if err != nil {
		return false, err.Error()
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return false, fmt.Sprintf("URL %s returned status code %d", url, response.StatusCode)
	}
	return true, ""
}
func processUrl(url string) {
	// Check if the URL is in valid format or not.
	valid, err := checkUrlValidity(url)
	if err != nil {
		fmt.Printf("There is a problem with proccessing URLs.")
	}
	if !valid {
		fmt.Printf("The provided URL [[%s]] is not in not a valid URL => Correct format: [http://example.com][https://example.com] \n", url)
		return
	}
	// Check if the URL is reachable or not.
	reachable, httpErr := checkUrlAvailability(url)
	if !reachable {
		fmt.Printf("%s is not reachable (error:%s)\n", url, httpErr)
		return
	}
	fmt.Printf("%s is reachable\n", url)

}
func main() {
	if len(os.Args) < 2 {
		fmt.Print("Usage: main.go URL1 URL2 ...")
		return
	}
	urls := os.Args[1:]
	wg := sync.WaitGroup{}
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			processUrl(url)
		}(url)
	}
	wg.Wait()

}
