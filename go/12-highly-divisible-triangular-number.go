package main

import "fmt"

func countDivisors(n int) int {
	numDivisors := 0
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			if i*i == n {
				numDivisors++
			} else {
				numDivisors += 2
			}
		}
	}
	return numDivisors
}

func divisibleTriangleNumber(n int) int {
	index := 1
	currentTriangleNumber := 1
	for {
		if countDivisors(currentTriangleNumber) > n {
			return currentTriangleNumber
		}
		index++
		currentTriangleNumber += index
	}
}

func main() {
	fmt.Println(divisibleTriangleNumber(500))
}
