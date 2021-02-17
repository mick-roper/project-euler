package problem1

func sumOfMultiplesOf3And5(ceiling int) int {
	naturalNumbers := []int{}
	for i := 1; i < ceiling; i++ {
		if i%3 == 0 || i%5 == 0 {
			naturalNumbers = append(naturalNumbers, i)
		}
	}

	sum := 0
	for _, val := range naturalNumbers {
		sum += val
	}
	return sum
}
