package main

import "fmt"

func main() {
 ifTest()
 switchTest()
}

func ifTest()  {
	intValue := 10

	//goland:noinspection GoBoolExpressions
	if intValue > 1 {
		println("hi")
	} else {
		println("ha")
	}
}

func switchTest() {
	intValue := 10

	switch intValue {
		case 1:
			println("1")
		case 2:
			println("2")
		case 3:
			println("3")
		case 4:
			println("4")
		case 8, 10:
			println("10")
			//goland:noinspection GoBoolExpressions
			if intValue == 10 {
				break
			}
			fallthrough
		default:
			println("default")
	}

	var undefinedValue interface{}

	switch i := undefinedValue.(type) {
		case nil:
			fmt.Println(" undefinedValue type :%T", i)
		case int:
			fmt.Println("undefinedValue type is int")
		case float64:
			fmt.Println("undefinedValue type is float64")
		case func(int) float64:
			fmt.Println("undefinedValue type is func(int)")
		case bool, string:
			fmt.Println("undefinedValue type is bool or string" )
		default:
			println("hi")
	}
}
