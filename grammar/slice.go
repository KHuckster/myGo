package main

import "fmt"

func main() {
	sliceTest()
}

func sliceTest()  {
	slice := make([]int, 10)
	printlnSlice(slice)
	//error
	//slice[10] = 20

	slice = make([]int, 20, 50)
	printlnSlice(slice)
	//error
	//slice[20] = 20

	intArray := [4]int{1, 2, 3, 4}
	slice = intArray[:]
	printlnSlice(slice)
	slice = intArray[0:2]
	printlnSlice(slice)
	slice = intArray[0:4]
	printlnSlice(slice)

	var nilSlice []int
	printlnSlice(nilSlice)
	//goland:noinspection GoNilness
	if nilSlice == nil {
		fmt.Println("nilSlice is null")
	} else {
		fmt.Println("nilSlice is not null")
	}

	slice = make([]int, 3)
	printlnSlice(slice)
	slice = append(slice, 1, 2, 3)
	printlnSlice(slice)

	slice2 := make([]int, len(slice), cap(slice))
	copy(slice2, slice)
	printlnSlice(slice2)
}

func printlnSlice(slice []int)  {
	fmt.Printf("slice len is %d, capcity is %d, slice is %x\n", len(slice), cap(slice), slice)
}
