package solution6

func sumSquareDifference(ceiling int) int {
	sqSum := 0
	sumSq := 0

	for i := 1; i <= ceiling; i++ {
		sumSq += i * i
		sqSum += i
	}

	sqSum = sqSum * sqSum
	return sqSum - sumSq
}
