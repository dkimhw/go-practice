package concurrency

type WebsiteChecker func(string) bool

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	for _, url := range urls {
		/*
			Firstly, they can be executed at the same time that they're declared - this is
			what the () at the end of the anonymous function is doing.

			Secondly they maintain access to the lexical scope in which they are defined
			- all the variables that are available at the point when you declare the anonymous
			function are also available in the body of the function.
		*/
		go func() {
			results[url] = wc(url)
		}()
	}

	return results
}
