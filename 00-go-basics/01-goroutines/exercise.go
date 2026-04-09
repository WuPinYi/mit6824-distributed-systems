package main

import (
	"fmt"
	"time"
)

// =============================================================
// EXERCISE: Goroutines
// =============================================================
// Instructions: Read each TODO and fill in the code.
// Run with: go run exercise.go
// =============================================================

// countDown counts down from n to 1, printing each number.
// It should print: "[name]: 3", "[name]: 2", "[name]: 1", "[name]: Done!"
func countDown(name string, n int) {
	for i := n; i >= 1; i-- {
		fmt.Printf("%s: %d\n", name, i)
		time.Sleep(200 * time.Millisecond) // simulate work
	}
	fmt.Printf("%s: Done!\n", name)
}

func main() {
	fmt.Println("--- Exercise 1: Basic Goroutine ---")
	// TODO 1: Call countDown("Alice", 3) as a GOROUTINE (runs in background)
	// Hint: just add "go" in front of the function call
	go countDown("Alice", 3) // change this line!

	// TODO 2: Also call countDown("Bob", 3) as a GOROUTINE
	// Your code here:
	go countDown("Bob", 3)

	// TODO 3: We need to wait for both goroutines to finish.
	// For now, use time.Sleep(2 * time.Second) to wait.
	// (We'll learn a better way in 03-mutex with WaitGroup!)
	// Your code here:
	time.Sleep(2 * time.Second)

	fmt.Println("--- Exercise 2: The Closure Bug ---")
	// TODO 4: The code below has the classic closure bug.
	// Fix it so it prints 0, 1, 2 (in any order) instead of 2, 2, 2
	for i := 0; i < 3; i++ {
		go func(number int) {
			fmt.Println("Number:", number) // BUG: fix this!
		}(i)
	}
	time.Sleep(500 * time.Millisecond) // wait for goroutines

	fmt.Println("All done!")
}

// =============================================================
// Expected output (order of Alice/Bob lines may vary):
//
// --- Exercise 1: Basic Goroutine ---
// Alice: 3
// Bob: 3
// Alice: 2
// Bob: 2
// Alice: 1
// Bob: 1
// Alice: Done!
// Bob: Done!
// --- Exercise 2: The Closure Bug ---
// Number: 0
// Number: 1
// Number: 2
// All done!
// =============================================================
