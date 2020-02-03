package main

import "fmt"

func helper(number int) string {
	singleDigits := []string{
		"", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
	}
	twoDigits := []string{
		"ten", "eleven", "twelve", "thirteen", "fourteen",
		"fifteen", "sixteen", "seventeen", "eighteen", "nineteen",
	}

	tens := []string{
		"", "ten", "twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety",
	}

	if number == 0 {
		return ""
	} else if number < 10 {
		return singleDigits[number]
	} else if number < 20 {
		return twoDigits[number-10]
	} else if number < 100 {
		return tens[number/10] + helper(number%10)
	} else {
		temp := singleDigits[number/100] + "hundred"
		if number%100 != 0 {
			temp += "and" + helper(number%100)
		}
		return temp
	}
}

func numberToWord(number int) string {
	thousands := []string{"", "thousand"}

	result := ""
	i := 0
	for number > 0 {
		if number%1000 != 0 {
			result += helper(number%1000) + thousands[i] + result
		}
		number /= 1000
		i++
	}

	return result
}

func numberLetterCounts(limit int) int {
	result := 0
	for i := 1; i <= limit; i++ {
		strNumber := numberToWord(i)
		result += len(strNumber)
	}

	return result
}

func main() {
	fmt.Println(numberLetterCounts(1000))
}
