package main

import "fmt"
import "math"
import "strconv"

func isPalindrome(n int) bool {
	s := strconv.Itoa(n)
	slen := len(s)

	i := 0
	j := slen - 1
	for i <= j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}

	return true
}

// Find the largest palindrome made fromthe product of two n-digit numbers
func largestPalindromeProduct(n int) int {
	start := int(math.Pow(10, float64(n-1)))
	end := int(math.Pow(10, float64(n)) - 1)

	var currProduct int
	maxProduct := 0
	for i := start; i <= end; i++ {
		for j := start; j <= end; j++ {
			currProduct = i * j
			if currProduct > maxProduct && isPalindrome(currProduct) {
				maxProduct = i * j
			}
		}

	}

	return maxProduct
}

func main() {
	fmt.Println(largestPalindromeProduct(3))
}
