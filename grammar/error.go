package main

import (
	"errors"
	"fmt"
)

type SelfError struct{}

func main() {
	var err error
	err = errorTest(0)
	if err != nil {
		fmt.Println(err)
	}

	err = errorTest(1)
	if err != nil {
		fmt.Println(err)
	}

	err = new(SelfError)
	fmt.Println(err.Error())
}

func errorTest(int1 int) error {
	if int1 == 0 {
		return errors.New("int1 is zero")
	}
	return nil
}

func (selfError SelfError) Error() string {
	return "i am SelfError"
}
