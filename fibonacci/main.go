package main

import (
	"fmt"
	"time"
)

var memo = make(map[int]int)

func shitFibonacci(n int) int {
	if n <= 2 {
		return 1
	} else if n > 40 {
		return 404
	}
	return shitFibonacci(n-1) + shitFibonacci(n-2)
}

func niceFibonacci(n int) int {
	// This uses memoization

	if memo[n] != 0 {
		return memo[n]
	} else if n <= 2 {
		return 1
	}

	memo[n] = (niceFibonacci(n-1) + niceFibonacci(n-2))
	return memo[n]
}

func godFibonacci(n int) int {
	// This uses memoization

	if n <= 2 {
		return 1
	}

	if memo[n] != 0 {
		return memo[n]
	}

	memo[1], memo[2] = 1, 1
	for i := 3; i <= n; i++ {
		memo[i] = memo[i-1] + memo[i-2]
	}

	return memo[n]
}

func main() {
	var n int

	// Get user input
	fmt.Print("Type the number of Fibonacci you want: ")
	fmt.Scan(&n)

	// Start timer
	startTime := time.Now()

	// Call shitFib() and stop the timer
	fmt.Printf("\n[shitFib] | n(%v) | Value: %v ; Time: %v", n, shitFibonacci(n), time.Since(startTime))

	// Restart the timer
	startTime = time.Now()

	// Call niceFib() and stop the timer
	fmt.Printf("\n[niceFib] | n(%v) | Value: %v ; Time: %v", n, niceFibonacci(n), time.Since(startTime))

	// Restart the timer and memo
	memo = make(map[int]int)
	startTime2 := time.Now()

	// Call godFib() and stop the timer
	fmt.Printf("\n[godFib]  | n(%v) | Value: %v ; Time: %v", n, godFibonacci(n), time.Since(startTime2))
}
