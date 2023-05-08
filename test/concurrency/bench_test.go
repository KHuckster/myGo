package concurrency

import (
	"testing"
	"time"
)

func slowWebsiteChecker(_ string) bool {
	time.Sleep(2 * 1000)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)

	for i := 0; i < len(urls); i++ {
		urls[i] = "url"
	}
	for i := 0; i <= b.N; i++ {
		CheckWebsites(slowWebsiteChecker, urls)
	}
}
