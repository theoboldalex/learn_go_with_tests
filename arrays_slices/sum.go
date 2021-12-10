package sum

func Sum(numbers []int) (sum int) {
	for _, num := range numbers {
		sum += num
	}

	return
}

func SumAll(numbers ...[]int) (sum []int) {
	for _, num := range numbers {
		sum = append(sum, Sum(num))
	}

	return
}

func SumAllTails(numbers ...[]int) (sum []int) {
	for _, num := range numbers {
		if len(num) == 0 {
			sum = append(sum, 0)
			continue
		}
		sum = append(sum, Sum(num[1:]))
	}

	return
}
