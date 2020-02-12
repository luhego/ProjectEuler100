package main

import "fmt"

// Helper function to find a value in a slice
func find(slice []int, val int) int {
	for i, item := range slice {
		if item == val {
			return i
		}
	}

	return -1
}

func isLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

// We calculate the first day for a given year. Monday -> 0, ..., Sunday -> 6
func getFirstDayOfYear(year int) int {
	if year == 1900 {
		return 0
	}

	currentYear := 1900
	firstDayOfYear := 0
	for currentYear < year {
		lastDayOfYear := firstDayOfYear
		if isLeapYear(currentYear) {
			lastDayOfYear++
		}

		firstDayOfYear = (lastDayOfYear + 1) % 7
		currentYear++
	}

	return firstDayOfYear
}

// Count the number of Sundays that fell on the first of the month
func countSundaysPerYear(year int) int {
	numSundays := 0

	daysPerMonth := []int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	leapYear := isLeapYear(year)
	numDays := 365
	if leapYear {
		numDays++
		daysPerMonth[2]++
	}

	day := 1
	// slice that contain the day number for each first of the month.
	// firstOfMonthDays = 1, 32, 60, 90, ...
	firstOfMonthDays := []int{1}
	for month := 1; month <= 11; month++ {
		day += daysPerMonth[month]
		firstOfMonthDays = append(firstOfMonthDays, day)
	}

	firstDayOfYear := getFirstDayOfYear(year)
	sunday := 1 + (6 - firstDayOfYear)

	for sunday <= numDays {
		if find(firstOfMonthDays, sunday) != -1 {
			numSundays++
		}
		sunday += 7
	}

	return numSundays
}

func countSundaysBetweenYears(firstYear int, lastYear int) int {
	startYear := firstYear
	result := 0
	for startYear <= lastYear {
		result += countSundaysPerYear(startYear)
		startYear++
	}

	return result
}

func main() {
	fmt.Println(countSundaysBetweenYears(1943, 1946))
	fmt.Println(countSundaysBetweenYears(1995, 2000))
	fmt.Println(countSundaysBetweenYears(1901, 2000))
}
