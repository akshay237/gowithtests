package goconcurrency

import (
	"reflect"
	"testing"
	"time"
)

const mockURL = "waat://furhurterwe.geds"

func mockWebsiteChecker(website string) bool {
	if website == mockURL {
		return false
	}
	return true
}

func TestCheckWebsite(t *testing.T) {
	websites := []string{
		"http://google.com",
		"http://blog.gypsydave5.com",
		"waat://furhurterwe.geds",
	}

	want := map[string]bool{
		"http://google.com":          true,
		"http://blog.gypsydave5.com": true,
		"waat://furhurterwe.geds":    false,
	}

	got := CheckWebsite(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got %v but want %v", got, want)
	}
}

func slowWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsite(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CheckWebsite(slowWebsiteChecker, urls)
	}
}

// To make it fast we will use concurrency
// Concurrency means having more than one thing in progress at the same time, that we do everyday in our day to day life.
// We say that this operation is blocking - it makes us wait for it to finish. An operation that does not block in Go will run in a separate process called a goroutine.
// we often use anonymous functions when we want to start a goroutine.
// An anonymous function literal looks just the same as a normal function declaration, but without a name.
// Anonymous functions have a number of features which make them useful, two of which we're using above.
// Firstly, they can be executed at the same time that they're declared - this is what the () at the end of the anonymous function is doing.
// Secondly they maintain access to the lexical scope in which they are defined - all the variables that are available at the point when you declare the anonymous function are also available in the body of the function.
// We can solve this data race by coordinating our goroutines using channels.
// Channels are a Go data structure that can both receive and send values.
// These operations, along with their details, allow communication between different processes.

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "some url"
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CheckWebsites(slowWebsiteChecker, urls)
	}
}
