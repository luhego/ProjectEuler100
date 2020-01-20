package main

import "fmt"

func isPrime(n int) bool {
	if n <= 3 {
		return n > 1
	} else if n%2 == 0 || n%3 == 0 {
		return false
	}

	i := 5
	for i*i <= n {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}

		i += 6
	}

	return true
}

func nthPrime(n int) int {
	i := 1
	primeCounter := 0

	for {
		if isPrime(i) {
			primeCounter++
		}

		if primeCounter == n {
			break
		}

		i++
	}
	return i
}

func main() {
	fmt.Println(nthPrime(1000000))
}
