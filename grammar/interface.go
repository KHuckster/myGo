package main

import "fmt"

type People interface {
	run()
	info()
}

type Man struct {
	age int
	name string
}

func main() {
	testInterface()
}

func (man Man)info(){
	fmt.Printf("man name is : %s, age is : %d\n", man.name, man.age)
}

func (man Man)run(){
	fmt.Println("go go")
}

func testInterface()  {
	man := Man{age: 18, name: "ZhangSan"}
	fmt.Println(man)
	man.info()
	man.run()

	var people People
	people = new(Man)
	people.info()
	people.run()

	people = man
	people.info()
	people.run()
}

