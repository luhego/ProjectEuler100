package main

import "fmt"

func getProperDivisionSum(n int) int {
	divisors := []int{1}

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			divisors = append(divisors, i)

			if i*i != n {
				divisors = append(divisors, n/i)
			}
		}
	}

	result := 0
	for i := 0; i < len(divisors); i++ {
		result += divisors[i]
	}

	return result
}

func sumAmicableNum(n int) int {
	divisorsSum := []int{0, 0}

	for i := 2; i <= n; i++ {
		divisorsSum = append(divisorsSum, getProperDivisionSum(i))
	}

	result := 0
	for i := 2; i <= n; i++ {
		if divisorsSum[i] <= n && i != divisorsSum[i] && i == divisorsSum[divisorsSum[i]] {
			result += i
		}
	}

	return result
}

func main() {
	fmt.Println(sumAmicableNum(10000))
}
