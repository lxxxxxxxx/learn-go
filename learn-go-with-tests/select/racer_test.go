package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T){
	// fastUrl := "www.baidu.com"
	// slowUrl:= "www.google.com"

	slowServer :=httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter,r *http.Request){
		time.Sleep(20*time.Millisecond)
		w.WriteHeader(http.StatusOK)
	}))

	fastServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter,r *http.Request){
		w.WriteHeader(http.StatusOK)
	}))

	slowUrl := slowServer.URL
	fastUrl := fastServer.URL

	want := fastUrl
	got := Racer(fastUrl,slowUrl)

	slowServer.Close()
	fastServer.Close()

	if got != want{
		t.Errorf("want %s but got %s",want,got)
	}


}
