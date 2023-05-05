package array

import (
	"reflect"
	"testing"
)

func TestSliceSum(t *testing.T) {
	returnCheck := func(t *testing.T, got, want int, numbers []int) {
		t.Helper()
		if got != want {
			t.Errorf("want %d but got %d, array %v", want, got, numbers)
		}
	}

	returnSumAllCheck := func(t *testing.T, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("want %v but got %v", want, got)
		}
	}

	t.Run("slice sum with iteration", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := sliceSum(numbers)
		want := 15
		returnCheck(t, got, want, numbers)
	})

	t.Run("slice sum all", func(t *testing.T) {
		got := sliceSumAll([]int{1, 2}, []int{2, 3})
		want := []int{3, 5}
		returnSumAllCheck(t, got, want)
	})

	t.Run("slice sum all append", func(t *testing.T) {
		got := sliceSumAllAppend([]int{1, 2}, []int{2, 3})
		want := []int{3, 5}
		returnSumAllCheck(t, got, want)
	})

	t.Run("slice sum with tails", func(t *testing.T) {
		got := sliceSumAllTails([]int{1, 2}, []int{2, 3}, []int{9})
		want := []int{2, 3, 0}
		returnSumAllCheck(t, got, want)
	})

	t.Run("empty slice sum with tails", func(t *testing.T) {
		got := sliceSumAllTails([]int{}, []int{2, 3}, []int{9})
		want := []int{0, 3, 0}
		returnSumAllCheck(t, got, want)
	})

}

func sliceSumAllTails(numbersSlice ...[]int) []int {
	var sumAll []int
	sumAll = make([]int, len(numbersSlice))
	for i, numbers := range numbersSlice {
		if len(numbers) == 0 {
			sumAll[i] = 0
			continue
		}
		sumAll[i] = sliceSum(numbers[1:])
	}
	return sumAll
}

func sliceSumAll(numbersSlice ...[]int) (sumAll []int) {
	sliceLen := len(numbersSlice)
	sumAll = make([]int, sliceLen)
	for i, numbers := range numbersSlice {
		sumAll[i] = sliceSum(numbers)
	}
	return
}

func sliceSumAllAppend(numbersSlice ...[]int) (sumAll []int) {
	for _, numbers := range numbersSlice {
		sumAll = append(sumAll, sliceSum(numbers))
	}
	return
}

func sliceSum(numbers []int) (sum int) {
	for _, number := range numbers {
		sum += number
	}
	return
}
