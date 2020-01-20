package main

import "fmt"

func specialPythagoreanTriplet(n int) int {
	sumOfabc := n

	for c := 3; c <= n; c++ {
		for b := 2; b < c; b++ {
			for a := 1; a < b; a++ {
				if a+b+c == sumOfabc && a*a+b*b == c*c {
					return a * b * c
				}
			}
		}
	}

	return -1
}

func main() {
	fmt.Println(specialPythagoreanTriplet(1000))
}
