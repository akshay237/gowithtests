package goconcurrency

import "time"

type WebsiteChecker func(string) bool

func CheckWebsite(wc WebsiteChecker, urls []string) map[string]bool {
	checkMap := map[string]bool{}

	for _, url := range urls {
		//checkMap[url] = wc(url)
		// to make it faster use go routine
		// if we use the url variable of the for loop then it will fail becz it will have only one url
		go func(url string) {
			checkMap[url] = wc(url)
		}(url)
	}
	time.Sleep(1 * time.Second)
	return checkMap
}

type resultChan struct {
	string
	bool
}

// Website Checker with channels
// We will be in safe side by using the channels because now no two go routines will write to the map at the same time.
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	checkMap := make(map[string]bool)
	checkChan := make(chan resultChan)

	for _, url := range urls {
		go func(u string) {
			checkChan <- resultChan{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		result := <-checkChan
		checkMap[result.string] = result.bool
	}

	return checkMap
}
