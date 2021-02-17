package problem3

import "math"

func largestPrimeFactor(candidate int) int {
	nPrime := 2
	x := candidate
	factors := []int{}

	if isPrime(candidate) {
		return candidate
	}

	for {
		// 1: check start prime / nprime is a whole number..
		if mod := math.Mod(float64(x), float64(nPrime)); mod == 0 {
			factor := x / nPrime
			factors = append(factors, nPrime)

			// 2: check if the result is prime...
			if isPrime(factor) {
				factors = append(factors, factor)
				break
			} else {
				x = factor
			}
		} else { // bump to the next prime
			nPrime = nextPrime(nPrime)
		}
	}

	xFactor := 0
	for _, yFactor := range factors {
		if yFactor > xFactor {
			xFactor = yFactor
		}
	}
	return xFactor
}

func isPrime(n int) bool {
	// handle first corner case
	if n <= 1 {
		return false
	}

	// handle second corner case
	if n <= 3 {
		return true
	}

	// this avoids having to check the middle 5 numbers in the following loop
	if n%2 == 0 || n%3 == 0 {
		return false
	}

	for i := 5; i*i < n; i = i + 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}

	return true
}

func nextPrime(n int) int {
	// base case
	if n <= 1 {
		return 2
	}

	prime := n

	for {
		prime++
		if isPrime(prime) {
			return prime
		}
	}
}
