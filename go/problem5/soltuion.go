package problem5

import "math"

func smallestMultiple(start, end int) int {
candidateLoop:
	for candidate := 1; ; candidate++ {
		for x := start; x <= end; x++ {
			if math.Mod(float64(candidate), float64(x)) != 0 {
				continue candidateLoop
			}
		}

		return candidate
	}
}
