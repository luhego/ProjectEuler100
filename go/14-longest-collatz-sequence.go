package main

import "fmt"

func longestCollatzSequence(limit int) int {
	maxLen := 0
	maxNumber := 0

	limit--
	for limit > 0 {
		currNumber := limit
		currLen := 1
		for currNumber != 1 {
			if currNumber%2 == 0 {
				currNumber /= 2
			} else {
				currNumber = 3*currNumber + 1
			}
			currLen++
		}

		if currLen > maxLen {
			maxLen = currLen
			maxNumber = limit
		}

		limit--
	}

	return maxNumber
}

func main() {
	fmt.Println(longestCollatzSequence(1000000))
}
