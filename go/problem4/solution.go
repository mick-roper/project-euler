package problem4

import (
	"fmt"
	"math"
)

func generatePalindrome(digits int) int {
	min := int(math.Pow10(digits - 1))
	max := int(math.Pow10(digits))
	var prod int

	for x := max - 1; x > min; x -= 2 {
		for y := max - 1; y > x; y -= 2 {
			prod = x * y
			if isPalindrome(prod) {
				return prod
			}
		}
	}

	panic("could not find a palindrome!")
}

func isPalindrome(product int) bool {
	str := fmt.Sprint(product)
	c := -1

	if len(str)%2 == 0 {
		c = len(str) / 2
	} else {
		c = len(str)/2 + 1
	}

	front := str[:c]
	rear := str[c:]

	return front == reverse(rear)
}

func reverse(str string) string {
	runes := []rune(str)
	res := make([]rune, len(str))
	j := 0
	i := len(runes) - 1
	for i >= 0 {
		res[i] = runes[j]
		j++
		i--
	}

	return string(res)
}
