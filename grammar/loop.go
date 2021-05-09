package main

func main() {
	forTest()
	forRangeTest()
	breakTest()
	continueTest()
	gotoTest()
}

func forTest()  {
	sum := 0

	for sum < 10 {
		sum ++
		println(sum)
	}
}

func forRangeTest()  {
	println("range test")

	strings := []string {"hello", "world"}

	for i, s := range strings {
		println(i, s)
	}
}

func breakTest() {
	println("break test")

	intArray := []int {1, 2, 3, 6, 10, 4, 14}

	for idx, element := range intArray {
		if element == 10 {
			break
		}
		println(idx, element)
	}

	exit:
		for _, element := range intArray {
			for idx1 := range intArray {
				if idx1 == 3 {
					break exit
				}
				println(idx1, element)
			}
		}
}

func continueTest()  {
	println("continue test")

	intArray := []int {1, 2, 3, 6, 10, 4, 14}

	for idx, element := range intArray {
		if element == 10 {
			continue
		}
		println(idx, element)
	}

	continuePoint:
		for _, element := range intArray {
			for idx1 := range intArray {
				if idx1 == 3 {
					continue continuePoint
				}
				println(idx1, element)
			}
		}
}

func gotoTest()  {
	println("goto test")

	intArray := []int {1, 2, 3, 6, 10, 4, 14}

	int1, int2 := 8, 9

	println(int1)

	gotoPoint:
		println(int2)

	for idx, element := range intArray {
		if element == 10 {
			goto gotoPoint
		}
		println(idx, element)
	}
}
