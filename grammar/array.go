package main

import "fmt"

func main() {
	arrayTest()

	intArray := [...]int{1: 10, 6: 234, 7: 34}
	arrayArgTest(intArray)
}

func arrayTest() {
	intArray := [3]int{67, 234, 34}
	fmt.Println(intArray)

	intArray2 := [...]int{1: 10, 6: 234, 10: 34}
	fmt.Println(intArray2)
	intArray2[3] = 14
	fmt.Println(intArray2[3])

	for index, element := range intArray2 {
		fmt.Println(index, element)
	}
	fmt.Println(len(intArray2))
}

func arrayArgTest(intArray [8]int) {
	for index, element := range intArray {
		fmt.Println(index, element)
	}
}
