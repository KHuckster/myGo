package main

type cb func(a int)

type person struct {
	name string
}

func main() {
	println("test swap")
	int1, int2 := 1, 2
	println(int1, int2)
	valueSwap(int1, int2)
	println(int1, int2)
	pointSwap(&int1, &int2)
	println(int1, int2)

	funcVar := testFuncVariable
	println("test funcVariable")
	println(funcVar(2))
	println(funcVar)
	println(&funcVar)

	println("test funcVariable with init arg")
	funcVar2 := testFuncVariable(2)
	println(funcVar2)
	println(&funcVar2)

	println("test callback")
	testCallBack(9, callback)
	testCallBack(9, func(a int) {
		println("hello world")
	})
	testCallBack2(10, callback)
	testCallBack2(10, func(a int) {
		println("hello world")
	})

	println("test closure function")
	sequence := getSequence()
	println(sequence())
	println(sequence())
	println(sequence())

	sequence1 := getSequence()
	println(sequence1())
	println(sequence1())
	println(sequence1())

	println("test closure function with init arg")
	sequenceWithArg1 := getSequenceWithArg(0)
	println(sequenceWithArg1(0))
	println(sequenceWithArg1(0))
	println(sequenceWithArg1(0))

	sequenceWithArg2 := getSequenceWithArg(0)
	println(sequenceWithArg2(0))
	println(sequenceWithArg2(1))
	println(sequenceWithArg1(2))

	sequenceWithArg3 := getSequenceWithArg
	println(sequenceWithArg3(1)(0))
	println(sequenceWithArg3(1)(1))
	println(sequenceWithArg3(1)(2))

	sequenceWithArg4 := getSequenceWithArg
	println(sequenceWithArg4(1))
	println(sequenceWithArg4(2))
	println(sequenceWithArg4(3))

	println("test method")
	var p person
	println(p.getName())
	p.name = "zhang san"
	println(p.getName())
	p.rename()
	println(p.getName())
	p.rightRename()
	println(p.getName())
	changeName(p, "zhang san")
	println(p.getName())
	realChangeName(&p, "zhang san")
	println(p.getName())
	println(&p)
}

func valueSwap(a, b int) {
	a, b = b, a
}

func pointSwap(a, b *int)  {
	var temp int
	temp = *a
	*a = *b
	*b = temp
	//*a, *b := *b, *a
}

func testFuncVariable(a int) int{
	println("hello", a)
	return a
}

func testCallBack(a int, callbackFunc cb)  {
	callbackFunc(a)
}

func testCallBack2(a int, callbackFunc func(a int))  {
	callbackFunc(a)
}

func callback(a int)  {
	println(a)
}

func getSequence() func() int{
	i := 0
	return func() int{
		i++
		return i
	}
}

func getSequenceWithArg(initInt int) func(int1 int) int{
	i := initInt
	return func(int1 int) int{
		i++
		return i + int1
	}
}

func (p person) getName() string{
	return p.name
}

func (p person) rename() {
	p.name = "li si"
}

func (p *person) rightRename() {
	p.name = "li si"
}

func changeName(p person, newName string) {
	p.name = newName
}

func realChangeName(p *person, newName string) {
	println(p)
	p.name = newName
}
