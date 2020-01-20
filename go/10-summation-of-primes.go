package main

import "fmt"

func primeSummation(n int) int {
	// We build the sieve of Eratosthenes
	// If sieve[i] = true, i is a prime number
	sieve := make([]bool, n+1)
	for i := 0; i <= n; i++ {
		sieve[i] = true
	}

	for p := 2; p*p <= n; p++ {
		if sieve[p] == true {
			for i := p * p; i <= n; i += p {
				sieve[i] = false
			}
		}
	}

	result := 0
	for i := 2; i < n; i++ {
		if sieve[i] == true {
			result += i
		}
	}
	return result
}

func main() {
	fmt.Println(primeSummation(2000000))
}
