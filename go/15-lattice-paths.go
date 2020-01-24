package main

import "fmt"
import "strconv"

func countLatticePaths(startX int, startY int, goalX int, goalY int, cache map[string]int) int {
	if startX > goalX || startY > goalY {
		return 0
	}

	if startX == goalX && startY == goalY {
		return 1
	}

	key := strconv.Itoa(startX) + "|" + strconv.Itoa(startY)

	cachedValue, exists := cache[key]
	if exists == true {
		return cachedValue
	}

	numPaths := 0

	// Move to the right
	numPaths += countLatticePaths(startX+1, startY, goalX, goalY, cache)

	// Move down
	numPaths += countLatticePaths(startX, startY+1, goalX, goalY, cache)

	cache[key] = numPaths

	return numPaths
}

func latticePaths(gridSize int) int {
	cache := make(map[string]int)
	return countLatticePaths(0, 0, gridSize, gridSize, cache)
}

func main() {
	fmt.Println(latticePaths(20))
}
