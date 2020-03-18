package main

import "fmt"
import "strings"

func swap(permutation []int, i int, j int) {
	temp := permutation[i]
	permutation[i] = permutation[j]
	permutation[j] = temp
}

func reverse(permutation []int, i int, num int) int {
	start := i
	end := len(permutation) - 1
	swapIndex := -1
	swapValue := 10

	for start <= end {
		swap(permutation, start, end)

		if permutation[start] > num && permutation[start] < swapValue {
			swapIndex = start
			swapValue = permutation[start]
		}

		if permutation[end] > num && permutation[end] < swapValue {
			swapIndex = end
			swapValue = permutation[end]
		}

		start++
		end--
	}

	return swapIndex
}

func nextPermutation(permutation []int) {
	i := len(permutation) - 1
	for permutation[i-1] > permutation[i] {
		i--
	}

	swapIndex := reverse(permutation, i, permutation[i-1])
	if swapIndex != -1 {
		swap(permutation, i-1, swapIndex)
	}
}

func lexicographicPermutations(n int) string {
	permutation := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := 0; i < n; i++ {
		nextPermutation(permutation)
	}

	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(permutation)), ""), "[]")
}

func main() {
	fmt.Println(lexicographicPermutations(999999))
}
