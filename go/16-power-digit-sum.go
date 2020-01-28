package main

import "fmt"

func multiplyByTwo(number []int) []int {
	carry := 0
	result := make([]int, 0)

	for i := 0; i < len(number); i++ {
		temp := number[i]*2 + carry
		carry = temp / 10
		result = append(result, temp%10)
	}

	if carry > 0 {
		result = append(result, carry)
	}

	return result
}

func powerDigitSum(exponent int) int {
	number := make([]int, 0)
	number = append(number, 2)

	for i := 0; i < exponent-1; i++ {
		number = multiplyByTwo(number)
	}

	result := 0
	for i := 0; i < len(number); i++ {
		result += number[i]
	}

	return result
}

func main() {
	fmt.Println(powerDigitSum(1000))
}
