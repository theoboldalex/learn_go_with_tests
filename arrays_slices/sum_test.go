package sum

import "testing"

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
