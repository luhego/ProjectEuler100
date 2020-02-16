package main

import "io/ioutil"
import "fmt"
import "strings"
import "sort"

func getNameScore(name string) int {
	score := 0
	for i := 0; i < len(name); i++ {
		score += (int(name[i]) - 64)
	}

	return score
}

func namesScores(arr []string) int {
	sort.Strings(arr)

	score := 0
	for i := 0; i < len(arr); i++ {
		score += getNameScore(arr[i]) * (i + 1)
	}

	return score
}

func main() {
	data, _ := ioutil.ReadFile("p022_names.txt")

	test := strings.Split(string(data), ",")
	for i := 0; i < len(test); i++ {
		test[i] = strings.ReplaceAll(test[i], "\"", "")
	}

	fmt.Println(namesScores(test))
}
