package array

import "testing"

func TestSum(t *testing.T) {
	returnCheck := func(t *testing.T, got, want int, numbers [5]int) {
		t.Helper()
		if got != want {
			t.Errorf("want %d but got %d, array %v", want, got, numbers)
		}
	}
	numbers := [5]int{1, 2, 3, 4, 5}
	want := 15

	t.Run("sum with iteration", func(t *testing.T) {
		got := sum(numbers)
		returnCheck(t, got, want, numbers)
	})

	t.Run("sum with range", func(t *testing.T) {
		got := sumWithRange(numbers)
		returnCheck(t, got, want, numbers)
	})
}

func sumWithRange(numbers [5]int) (sum int) {
	for _, number := range numbers {
		sum += number
	}
	return
}

func sum(numbers [5]int) (sum int) {
	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}
	return
}
