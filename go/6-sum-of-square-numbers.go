package main

import "fmt"
import "math"

func sumSquareDifference(n int) int {
	sumOfSquares := (n * (n + 1) * (2*n + 1)) / 6
	sumN := (n * (n + 1)) / 2
	squareOfSum := int(math.Pow(float64(sumN), 2.0))

	return int(math.Abs(float64(sumOfSquares - squareOfSum)))
}

func main() {
	fmt.Println(sumSquareDifference(100))
}
