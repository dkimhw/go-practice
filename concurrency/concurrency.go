package concurrency

type WebsiteChecker func(string) bool
type result struct {
	// We can assign anonymous values within the struct - can
	// be useful in when it's hard to know what to name a value.
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		/*
			Firstly, they can be executed at the same time that they're declared - this is
			what the () at the end of the anonymous function is doing.

			Secondly they maintain access to the lexical scope in which they are defined
			- all the variables that are available at the point when you declare the anonymous
			function are also available in the body of the function.
		*/
		go func(u string) {
			// Send statement
			resultChannel <- result{u, wc(u)}
		}(url)
	}

	/*
		Using channels, we can control the timing of each write into the
		results map, ensuring that it happens one at a time. Although
		each of the calls of `wc` and each send to the result channel, is
		happening concurrently inside its own process, each of the results
		is being dealt with one at a time as we take values out of the result
		channel with the receive expression.
	*/
	for i := 0; i < len(urls); i++ {
		// Receive expression
		r := <-resultChannel
		results[r.string] = r.bool
	}

	return results
}
