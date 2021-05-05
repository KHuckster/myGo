package main

import "fmt"

var globalInt int = 2

var (
	globalString = "global hello"
	globalBoolean = "global bool"
	//globalInt2 := 3
)

func main() {
	printStringVariable()
	printIntVariable()
	printDefaultValue()
	printGlobalVariable()
	printSwapVariable()
	_, int2, string1 := returnVariable()
	fmt.Println(int2)
	fmt.Println(string1)

}

func printStringVariable() {
	var hello string = "hello"
	fmt.Println(hello)
}

func printIntVariable() {
	var int1, int2 int = 1, 2
	fmt.Println(int1)
	fmt.Println(int2)
}

func printDefaultValue() {
	var defaultString string
	var defaultInt int
	var defaultBoolean bool
	var defaultValue = "hello"
	fmt.Println(defaultString)
	fmt.Println(defaultInt)
	fmt.Println(defaultBoolean)
	fmt.Println(defaultValue)

	//var defaultInt2 int
	defaultInt2 := 2
	defaultString2 := "hello2"
	fmt.Println(defaultInt2)
	fmt.Println(defaultString2)
}

func printGlobalVariable()  {
	fmt.Println(globalInt)
	fmt.Println(globalString)
	globalString = "world"
	fmt.Println(globalString)
	fmt.Println(globalBoolean)
}

func printSwapVariable() {
	var int1, int2 = 2, 3
	int1, int2 = int2, int1
	fmt.Println(int1)
	fmt.Println(int2)
}

func returnVariable() (int, int, string) {
	return 1, 2, "hello"
}
