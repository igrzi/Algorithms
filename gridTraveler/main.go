package main

import (
	"fmt"
	"strconv"
	"time"
)

func shitGridTraveler(m int, n int) int {
	if m == 1 && n == 1 {
		return 1
	} else if m == 0 || n == 0 {
		return 0
	} else if m > 15 || n > 15 {
		return 404
	}
	return shitGridTraveler(m-1, n) + shitGridTraveler(m, n-1)
}

var memo = make(map[string]int)

func niceGridTraveler(m int, n int) int {
	key := strconv.Itoa(m) + "," + strconv.Itoa(n)

	if memo[key] != 0 {
		return memo[key]
	}

	if m == 1 && n == 1 {
		return 1
	} else if m == 0 || n == 0 {
		return 0
	}
	memo[key] = niceGridTraveler(m-1, n) + niceGridTraveler(m, n-1)
	return memo[key]
}

func godGridTraveler(m int, n int) int {
	// Create a 2D grid to store the number of ways to reach each cell
	grid := make([][]int, m+1)
	for i := range grid {
		grid[i] = make([]int, n+1)
	}

	// Initialize base cases
	grid[1][1] = 1

	// Fill the grid using dynamic programming
	for i := 0; i <= m; i++ {
		for j := 0; j <= n; j++ {
			if i+1 <= m {
				grid[i+1][j] += grid[i][j]
			}
			if j+1 <= n {
				grid[i][j+1] += grid[i][j]
			}
		}
	}

	// The result is in the bottom-right cell
	return grid[m][n]
}

func main() {
	var sizeX, sizeY int

	fmt.Print("Type the size of the grid, separated by a space: ")
	fmt.Scan(&sizeX, &sizeY)

	// Start timer
	startTime := time.Now()
	fmt.Printf("\n[shitGrid] | grid(%v, %v) | Value: %v | Time: %v", sizeX, sizeY, shitGridTraveler(sizeX, sizeY), time.Since(startTime))

	// Start timer
	startTime = time.Now()
	fmt.Printf("\n[niceGrid] | grid(%v, %v) | Value: %v | Time: %v", sizeX, sizeY, niceGridTraveler(sizeX, sizeY), time.Since(startTime))

	startTime = time.Now()
	fmt.Printf("\n[godGrid]  | grid(%v, %v) | Value: %v | Time: %v\n\n", sizeX, sizeY, godGridTraveler(sizeX, sizeY), time.Since(startTime))
}
