package main

import "fmt"

func fiboEvenSum() int {
	sum := 0
	a := 1
	b := 1
	var temp int
	for b <= 4000000 {
		temp = a
		a = b
		b = a + temp

		if b%2 == 0 {
			sum += b
		}
	}

	return sum
}

func main() {
	fmt.Println(fiboEvenSum())
}
