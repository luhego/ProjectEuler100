package main

import "fmt"

func gcd(a int, b int) int {
	if a == 0 {
		return b
	}
	return gcd(b%a, a)
}

func lcm(a int, b int) int {
	return (a * b) / gcd(a, b)
}

func smallestMult(n int) int {
	result := 1
	for i := 2; i <= n; i++ {
		result = lcm(result, i)
	}

	return result
}

func main() {
	fmt.Println(smallestMult(20))
}
