package main

import (
	"reflect"
	"testing"
)

type MockWebsitChecker struct{}

func (m MockWebsitChecker) WebsitChecker(url string) bool {
	if url == "www.abcd.com" {
		return false
	}
	return true
}

func BenchmarkConcurrency(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}
	mockWebsitChecker := MockWebsitChecker{}
	for i := 0; i < b.N; i++ {
		CheckWebsit(&mockWebsitChecker, urls)
	}
}

func TestConcurrency(t *testing.T) {
	urls := []string{
		"www.google.com",
		"www.baidu.com",
		"www.abcd.com",
	}
	mockWebsitChecker := MockWebsitChecker{}

	res := CheckWebsit(&mockWebsitChecker, urls)

	want := len(urls)
	got := len(res)

	if want != got {
		t.Errorf("want %d results,but got %d", want, got)
	}

	wantResult := map[string]bool{
		"www.google.com": true,
		"www.baidu.com":  true,
		"www.abcd.com":   false,
	}

	if !reflect.DeepEqual(res, wantResult) {
		t.Errorf("want %v ,but got %v", wantResult, res)
	}
}
