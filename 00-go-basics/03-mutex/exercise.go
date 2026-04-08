package main

import (
	"fmt"
	"sync"
)

// =============================================================
// EXERCISE: Mutex & WaitGroup
// =============================================================
// Instructions: Read each TODO and fill in the code.
// Run with:       go run exercise.go
// Check for races: go run -race exercise.go
// =============================================================

// SafeCounter is a counter that is safe for concurrent use
type SafeCounter struct {
	// TODO 1: Add a sync.Mutex field named "mu"
	// TODO 2: Add an int field named "count"
	// Your code here:
}

// Increment adds 1 to the counter safely
func (c *SafeCounter) Increment() {
	// TODO 3: Lock the mutex before touching c.count, unlock after
	// Your code here:
	c.count++ // this line is fine, just add lock/unlock around it
}

// Value returns the current count safely
func (c *SafeCounter) Value() int {
	// TODO 4: Lock the mutex before reading c.count, unlock after
	// Hint: use defer for the unlock!
	// Your code here:
	return c.count
}

func main() {
	fmt.Println("--- Exercise 1: Safe Counter with Mutex ---")

	counter := SafeCounter{}

	// TODO 5: Use a WaitGroup to launch 1000 goroutines that each call counter.Increment()
	// Then wait for all of them to finish, and print counter.Value()
	// Expected output: 1000
	//
	// Steps:
	//   a) Declare a sync.WaitGroup
	//   b) Loop 1000 times, each time: wg.Add(1) then launch goroutine
	//   c) Inside goroutine: defer wg.Done(), then call counter.Increment()
	//   d) After loop: wg.Wait()
	//   e) Print counter.Value()
	//
	// Your code here:

	fmt.Println("--- Exercise 2: Parallel Results Collection ---")

	// TODO 6: Launch 5 goroutines, each computing i*i (for i = 1 to 5)
	// Collect all results into a shared []int slice named "results"
	// Use a mutex to protect the slice when appending
	// Use a WaitGroup to wait for all goroutines
	// Print results at the end (order doesn't matter)
	//
	// Expected output (any order): [1 4 9 16 25]
	//
	// Your code here:

	fmt.Println("--- Exercise 3: Race Condition Detection ---")
	// This code has a race condition. Run it with:
	//   go run -race exercise.go
	// You should see a WARNING from the race detector.
	// Then fix it by adding a mutex.

	var shared int
	var wg2 sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg2.Add(1)
		go func() {
			defer wg2.Done()
			shared++ // ← RACE CONDITION! Fix this with a mutex.
		}()
	}
	wg2.Wait()
	fmt.Println("Shared value (with race condition):", shared) // might not be 100!
}
