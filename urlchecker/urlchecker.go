package urlchecker

import (
	"fmt"
	"net/http"
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

func checkUrlAvailability(url string) (bool, error) {
	response, err := http.Get(url)
	if err != nil {
		return false, err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return false, fmt.Errorf("%s returned status code %d", url, response.StatusCode)
	}
	return true, nil
}

func ProcessUrl(wg *sync.WaitGroup, url string) {
	defer wg.Done()
	// Check if the URL is in valid format or not.
	valid, err := checkUrlValidity(url)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	if !valid {
		fmt.Printf("The provided URL [[%s]] is not in not a valid URL => Correct format: [https://example.com] \n", url)
		return
	}
	// Check if the URL is reachable or not.
	reachable, err := checkUrlAvailability(url)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	if !reachable {
		fmt.Printf("%s is not reachable\n", url)
		return
	}
	fmt.Printf("%s is reachable\n", url)
}
