package iteration

import "testing"

func TestRepeat(t *testing.T) {
	returnCheck := func(t *testing.T, repeated string, expected string) {
		t.Helper()
		if repeated != expected {
			t.Errorf("expected '%q' but repeated '%q'", expected, repeated)
		}
	}
	expected := "aaaaa"

	t.Run("Repeat", func(t *testing.T) {
		returnCheck(t, Repeat("a"), expected)
	})

	t.Run("Repeat2", func(t *testing.T) {
		returnCheck(t, Repeat2("a"), expected)
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}

func Repeat(character string) (repeated string) {
	for i := 0; i < 5; i++ {
		repeated += character
	}
	return
}

func Repeat2(character string) string {
	var repeated string
	for i := 0; i < 5; i++ {
		repeated += character
	}
	return repeated
}
