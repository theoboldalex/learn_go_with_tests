package sum

func Sum(numbers []int) (sum int) {
	for _, num := range numbers {
		sum += num
	}

	return
}
