package main

import "unsafe"

const (
	ConstValue1 = "hello world!"
	ConstValue2 = len(ConstValue1)
	ConstValue3 = unsafe.Sizeof(ConstValue1)
	ConstValue4 = unsafe.Sizeof(ConstValue2)
)

const (
	IotaConstValue1 = iota
	IotaConstValue2
	IotaConstValue3 = 5 * iota
	IotaConstValue4
	IotaConstValue5 = 2 * iota
	IotaConstValue6
)

func main() {
	printConstantValue()
	printConstantIotaValue()
}

func printConstantValue() {
	const ConstValue = "hello"
	//ConstValue = "world"
	println(ConstValue)
	println(ConstValue1)
	println(ConstValue2)
	println(ConstValue3)
	println(ConstValue4)
}

func printConstantIotaValue() {
	println(IotaConstValue1)
	println(IotaConstValue2)
	println(IotaConstValue3)
	println(IotaConstValue4)
	println(IotaConstValue5)
	println(IotaConstValue6)
}