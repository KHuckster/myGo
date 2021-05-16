package main

import (
	"fmt"
	"time"
)

func main() {
	testGoroutine()
	time.Sleep(100 * time.Second)

}

func testGoroutine() {
	go printString("async run")
	printString("main thread run")
}

func printString(str string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Second)
		fmt.Println(str)
	}
}
