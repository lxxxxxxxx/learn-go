package main

import (
	"fmt"
)

type IWebsitChecker interface {
	WebsitChecker(url string) bool
}

type RealWebsitChecker struct{}

func (r *RealWebsitChecker) WebsitChecker(url string) bool {

	return true
}

type Result struct {
	string
	bool
}

func CheckWebsit(wc IWebsitChecker, urls []string) map[string]bool {
	res := make(map[string]bool)
	resChan := make(chan Result)

	for _, url := range urls {
		go func(u string) {
			resChan <- Result{u, wc.WebsitChecker(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		one := <-resChan
		res[one.string] = one.bool
	}
	return res
}

func main() {
	urls := []string{
		"www.google.com",
		"www.baidu.com",
		"www.abcd.com",
	}
	realWebsitChecker := RealWebsitChecker{}
	res := CheckWebsit(&realWebsitChecker, urls)

	fmt.Printf("res : %v", res)

}
