package main

import "fmt"

func largestPrimeFactor(number int) int {
	primeFactor := 2
	result := 2

	for primeFactor <= number {
		if primeFactor > result && number%primeFactor == 0 {
			result = primeFactor
		}
		for number%primeFactor == 0 {
			number /= primeFactor
		}
		primeFactor++
	}

	return result
}

func main() {
	fmt.Println(largestPrimeFactor(600851475143))
}
