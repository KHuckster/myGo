package dictionary

import (
	"fmt"
	"testing"
)

const key = "network"
const keyNil = "keyNil"
const value = "computer network"

func TestDictionary(t *testing.T) {
	t.Run("found key", func(t *testing.T) {
		dictionary := Dictionary{key: value}
		got, _ := dictionary.Search(key)
		checkValue(t, got, value)
	})

	t.Run("notfound key", func(t *testing.T) {
		dictionary := Dictionary{key: value}
		got, err := dictionary.Search(keyNil)
		if err == nil {
			checkValue(t, got, value)
		} else {
			fmt.Println("expected to get an error.", err)
		}
	})
}

func TestAdd(t *testing.T) {
	t.Run("add key", func(t *testing.T) {
		dictionary := make(Dictionary)
		dictionary.Add(keyNil, value)

		got, _ := dictionary.Search(keyNil)
		checkValue(t, got, value)
	})
}

func checkValue(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
