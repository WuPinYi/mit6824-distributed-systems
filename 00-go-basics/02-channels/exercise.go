package main

import (
	"fmt"
	"time"
)

// =============================================================
// EXERCISE: Channels
// =============================================================
// Instructions: Read each TODO and fill in the code.
// Run with: go run exercise.go
// =============================================================

// square returns the square of n via a channel
func square(n int, ch chan int) {
	time.Sleep(100 * time.Millisecond) // simulate work
	// TODO 1: Send n*n into the channel ch
	// Your code here:
	ch <- n * n
}

// pipeline: double takes values from "in", doubles them, sends to "out"
func double(in chan int, out chan int) {
	for v := range in {
		// TODO 4: Send v*2 into the "out" channel
		// Your code here:
		out <- v * 2
	}
	close(out) // close output when input is done
}

func main() {
	fmt.Println("--- Exercise 1: Basic Channel ---")
	// TODO 2: Create a channel of type int
	// var ch chan int = ???
	// Your code here:
	ch := make(chan int)

	// Launch goroutines to compute squares
	go square(3, ch)
	go square(5, ch)
	go square(7, ch)

	// TODO 3: Receive 3 results from the channel and print them
	// Hint: use a for loop that runs 3 times, receive from ch each time
	// Your code here:
	for i := 0; i < 3; i++ {
		results := <-ch
		fmt.Println("Square result:", results)
	}

	fmt.Println("--- Exercise 2: Done Channel Pattern ---")
	// TODO 5: Use a channel to wait for a goroutine (instead of time.Sleep)
	// Create a "done" channel of type bool
	// Launch a goroutine that prints "working..." then sleeps 300ms then sends true to done
	// Wait for the signal on done channel before printing "goroutine finished!"
	// Your code here:
	done := make(chan bool)
	go func() {
		fmt.Println("working...")
		time.Sleep(300 * time.Millisecond)
		done <- true
	}()
	<-done
	fmt.Println("goroutine finished!")

	fmt.Println("--- Exercise 3: Pipeline ---")
	// TODO 6: Create a pipeline:
	//   numbers channel  →  double()  →  results channel
	//
	// Step 1: Create "numbers" channel (buffered, capacity 5)
	// Step 2: Create "results" channel (buffered, capacity 5)
	// Step 3: Send numbers 1,2,3,4,5 into "numbers" channel, then close it
	// Step 4: Launch double() as a goroutine with numbers and results channels
	// Step 5: Range over "results" and print each value
	//
	// Expected output: 2, 4, 6, 8, 10
	// Your code here:
	numbers := make(chan int, 5)
	results := make(chan int, 5)

	for i := 1; i <= 5; i++ {
		numbers <- i
	}
	close(numbers)

	go double(numbers, results)

	for v := range results {
		fmt.Println("Doubled:", v)
	}

	fmt.Println("All done!")
}
