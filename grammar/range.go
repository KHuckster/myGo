package main

import "fmt"

func main() {
	stringRangeTest()
}

func stringRangeTest()  {
	for index, element := range "abcd" {
		fmt.Println(index, element)
	}
}
