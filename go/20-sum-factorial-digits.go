package main

import "fmt"

func multiply(number []int, x int) []int {
	result := []int{}
	carry := 0

	for i := 0; i < len(number); i++ {
		temp := number[i]*x + carry
		carry = temp / 10
		result = append(result, temp%10)
	}

	for carry > 0 {
		result = append(result, carry%10)
		carry = carry / 10
	}

	return result
}

func sumFactorialDigits(n int) int {
	number := []int{1}

	for i := 1; i <= n; i++ {
		number = multiply(number, i)
	}

	sum := 0
	for i := 0; i < len(number); i++ {
		sum += number[i]
	}

	return sum
}

func main() {
	fmt.Println(sumFactorialDigits(100))
}
