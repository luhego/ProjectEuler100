package main

import "fmt"
import "sort"

func isAbundantNumber(n int) bool {
	proverDivisorsSum := 1

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			proverDivisorsSum += i

			if i*i != n {
				proverDivisorsSum += n / i
			}
		}
	}

	return proverDivisorsSum > n
}

func sumOfNonAbundantNumbers(n int) int {
	abundantNumbers := make(map[int]bool)
	for i := 1; i <= n; i++ {
		if isAbundantNumber(i) {
			abundantNumbers[i] = true
		}
	}

	alreadyIncludedNumber := make(map[int]bool)
	sumOfAbundantNumbers := make([]int, 0)
	for i := range abundantNumbers {
		for j := range abundantNumbers {
			sum := i + j
			if sum > n {
				continue
			}

			_, exists := alreadyIncludedNumber[sum]
			if exists == true {
				continue
			}
			sumOfAbundantNumbers = append(sumOfAbundantNumbers, sum)
			alreadyIncludedNumber[sum] = true
		}
	}

	sort.Ints(sumOfAbundantNumbers)

	result := 0
	current := 1
	for _, s := range sumOfAbundantNumbers {
		result += ((s - 1) * s) / 2
		if current > 1 {
			result -= ((current - 1) * current) / 2
		}

		current = s + 1
	}
	return result
}

func main() {
	fmt.Println(sumOfNonAbundantNumbers(28123))
}
