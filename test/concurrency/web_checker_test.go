package concurrency

import (
	"reflect"
	"testing"
)

func mockWebsiteChecker(url string) bool {
	if url == "www.baidu.com" {
		return true
	}
	return false
}

func TestWebCheck(t *testing.T) {
	websites := []string{
		"www.baidu.com",
		"www.google.com",
		"www.alibaba.com",
	}

	checkResult := CheckWebsites(mockWebsiteChecker, websites)

	got := len(checkResult)
	want := len(websites)
	if got != want {
		t.Fatalf("Wanted %v, got %v", want, got)
	}

	expectedMap := map[string]bool{
		"www.baidu.com":   true,
		"www.google.com":  false,
		"www.alibaba.com": false,
	}
	if !reflect.DeepEqual(checkResult, expectedMap) {
		t.Fatalf("Wanted %v, got %v", expectedMap, checkResult)
	}

}
