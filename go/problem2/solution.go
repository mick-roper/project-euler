package problem2

func fibonacciSum(ceiling int) int {
	a := 1
	b := 2
	c := -1

	fibonacciNumbers := []int{a, b}

	for {
		c = a + b
		a = b
		b = c

		if c >= ceiling {
			break
		}

		fibonacciNumbers = append(fibonacciNumbers, c)
	}

	sum := 0

	for _, x := range fibonacciNumbers {
		if x%2 == 0 {
			sum += x
		}
	}

	return sum
}
