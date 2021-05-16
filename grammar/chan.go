package main

import (
	"fmt"
	"time"
)

var (
	sleep = sleepOneSecond
)

func main() {
	testChan()
	time.Sleep(3 * time.Second)
	testChanWithBuffer()
	time.Sleep(3 * time.Second)
	testChanRange()
}

func testChan() {
	c := make(chan int)
	intArr := []int{1, 2, 3, 4, 5, 6}
	go testTransferData(intArr[:(len(intArr)/2)], c)
	go testTransferData(intArr[(len(intArr)/2):], c)
	sleep()
	fmt.Println(<-c)
	sleep()
	fmt.Println(<-c)
	sleep()
}

func testChanWithBuffer() {
	c := make(chan int, 2)
	intArr := []int{1, 2, 3, 4, 5, 6}
	go testTransferData(intArr[:(len(intArr)/2)], c)
	go testTransferData(intArr[(len(intArr)/2):], c)
	go testTransferData(intArr, c)
	sleep()
	fmt.Println(<-c)
	sleep()
	fmt.Println(<-c)
	sleep()
	fmt.Println(<-c)
	sleep()
}

func testChanRange() {
	c := make(chan int, 2)
	intArr := []int{1, 2, 3, 4, 5, 6}
	go testTransferData(intArr[:(len(intArr)/2)], c)
	go testTransferData(intArr[(len(intArr)/2):], c)
	go testTransferData(intArr, c)
	go delayCloseChannel(c)
	for i := range c {
		fmt.Println(i)
	}
	fmt.Println("exist test chan range")
}

func testTransferData(intArr []int, c chan int) {
	sum := 0
	for _, element := range intArr {
		sum += element
	}
	c <- sum
	fmt.Printf("send %d to channel\n", sum)
}

func sleepOneSecond() {
	time.Sleep(1 * time.Second)
}

func delayCloseChannel(c chan int) {
	time.Sleep(6 * time.Second)
	fmt.Println("close chan")
	close(c)
}
