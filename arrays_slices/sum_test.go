package sum

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("sum a collection of any length", func(t *testing.T) {
		numbers := []int{2, 4, 6, 8}

		result := Sum(numbers)
		expected := 20

		if result != expected {
			t.Errorf("Expected '%d' but got '%d', Given '%v'", expected, result, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	result := SumAll([]int{1, 2}, []int{3, 4})
	expected := []int{3, 7}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected '%v' but got '%v'.", expected, result)
	}
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t testing.TB, result, expected []int) {
		t.Helper()
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected '%v' but got '%v'.", expected, result)
		}

	}
	t.Run("sum the tails of some slices", func(t *testing.T) {
		result := SumAllTails([]int{1, 2, 3}, []int{2, 3, 4, 5})
		expected := []int{5, 12}

		checkSums(t, result, expected)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		result := SumAllTails([]int{}, []int{2, 3, 4, 5})
		expected := []int{0, 12}

		checkSums(t, result, expected)
	})
}
