package calculator

func Add(nub ...int) int {
	total := 0
	for _, value := range nub {
		total += value
	}

	return total
}
