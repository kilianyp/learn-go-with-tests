package concurrency

type WebsiteChecker func(string) bool

type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	result_channel := make(chan result)

	for _, url := range urls {
		go func(u string) {
			result_channel <- result{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		result := <-result_channel
		results[result.string] = result.bool
	}
	return results
}
