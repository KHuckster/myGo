package main

import "fmt"

const ArraySize = 3

func main() {
	ptrTest()
	arrayPtrTest()
	ptrPtrTest()
}

func ptrTest()  {
	var intPtr *int
	fmt.Printf("intPtr address is %X\n", intPtr)
	if intPtr == nil {
		fmt.Println("intPtr is null")
	} else {
		fmt.Printf("intPtr is %d\n", *intPtr)
	}

	intValue := 10
	intPtr = &intValue
	fmt.Printf("intPtr address is %X\n", intPtr)
	if intPtr == nil {
		fmt.Println("intPtr is null")
	} else {
		fmt.Printf("intPtr is %d\n", *intPtr)
	}
}

func arrayPtrTest()  {
	intArray := [ArraySize]int{8, 9, 0}
	var intArrayPtr [ArraySize] *int

	for index, element := range intArray {
		intArrayPtr[index] = &element
	}

	for index, element := range intArrayPtr {
		fmt.Println(index, *element)
	}
}

func ptrPtrTest()  {
	//var intPtr *int

	intValue := 10
	intPtr := &intValue
	intPtrPtr := &intPtr
	fmt.Printf("intPtr address is %X\n", &intPtr)
	if intPtr == nil {
		fmt.Println("intPtr is null")
	} else {
		fmt.Printf("intPtr point value is %d\n", *intPtr)
	}
	fmt.Printf("intPtrPtr address is %X\n", &intPtrPtr)
	fmt.Printf("intPtrPtr point address is %X\n", *intPtrPtr)
	if intPtrPtr == nil {
		fmt.Println("intPtrPtr is null")
	} else {
		fmt.Printf("intPtrPtr point value is %d\n", **intPtrPtr)
	}
}
