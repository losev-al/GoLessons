package concurrency

// WebsiteChecker describe the method for check web site
type WebsiteChecker func(string) bool

type result struct {
	url         string
	checkResult bool
}

// CheckWebsites test websites by WebsiteChecker
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result, 5)
	for _, url := range urls {
		go func(u string, c chan result) {
			resultChannel <- result{u, wc(u)}
		}(url, resultChannel)
	}

	finishChannel := make(chan bool)
	go func(c chan result, f chan bool) {
		for i := 0; i < len(urls); i++ {
			result := <-c
			results[result.url] = result.checkResult
		}
		f <- true
	}(resultChannel, finishChannel)
	<-finishChannel
	return results
}
