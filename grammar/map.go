package main

import "fmt"

func main() {
	mapTest()
	mapDeleteTest()
}

func mapTest() {
	var cityMap map[string]string
	//cityMap ["Chengdu"] = "Sichuan"
	cityMap = make(map[string]string)
	cityMap["Chengdu"] = "Sichuan"
	cityMap["Shijiazhuang"] = "Hebei"
	cityMap["Beijing"] = "Beijing"

	for index, element := range cityMap {
		fmt.Println(index, element)
	}
	fmt.Println(cityMap["Chengdu"])
	fmt.Println(cityMap["chengdu"])
	value, ok := cityMap["chengdu"]
	if ok {
		fmt.Printf("chengdu exist %s\n", value)
	} else {
		fmt.Println("chengdu not exist")
	}

	fmt.Println(cityMap["Chengdu"])
	value2, ok := cityMap["Chengdu"]
	if ok {
		fmt.Printf("Chengdu exist %s\n", value2)
	} else {
		fmt.Println("Chengdu not exist")
	}
}

func mapDeleteTest() {
	cityMap := map[string]string{"Chengdu": "Sichuan", "Shijiazhuang": "Hebei", "Beijing" : "Beijing"}
	for index, element := range cityMap {
		fmt.Println(index, element)
	}
	delete(cityMap, "Shenyang")
	fmt.Println(cityMap["Chengdu"])
	delete(cityMap, "Chengdu")
	fmt.Println(cityMap["Chengdu"])
}
