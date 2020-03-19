package main

import "fmt"
import "math"

func countDigitsOfNthFibonacci(n int) int {
	sqrt5 := math.Sqrt(5)
	phi := (1 + sqrt5) / 2.0
	return 1 + int(math.Floor(float64(n)*math.Log10(phi)-math.Log10(sqrt5)))
}

func digitFibonacci(n int) int {
	i := 1

	for {
		numDigits := countDigitsOfNthFibonacci(i)
		if numDigits == n {
			break
		}
		i++
	}

	return i
}

func main() {
	fmt.Println(digitFibonacci(1000))
}
