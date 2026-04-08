# Go Basics for Distributed Systems 🐒

Before touching the labs, you need to understand these 3 superpowers in Go:

| # | Concept | Real-world analogy |
|---|---------|-------------------|
| 1 | **Goroutines** | Hiring extra workers to do jobs in parallel |
| 2 | **Channels** | A pipe that workers use to pass results to each other |
| 3 | **Mutex** | A "ONE AT A TIME" sign on a shared bathroom |

Master these 3 and you can handle any lab.

---

## Why do we need these?

In distributed systems, many things happen **at the same time**:
- Server A is processing a request
- Server B is sending a heartbeat
- Server C just crashed and needs to be detected

Go makes it easy to handle all of this with goroutines + channels + mutexes.

---

## Learning Order

1. 👉 [01-goroutines](./01-goroutines/) — Start here
2. 👉 [02-channels](./02-channels/) — After goroutines
3. 👉 [03-mutex](./03-mutex/) — Last piece of the puzzle

---

## Quick Cheat Sheet

```go
// Goroutine — run something in background
go doSomething()

// Channel — make a pipe, send and receive
ch := make(chan int)
ch <- 42        // send 42 into the pipe
value := <-ch   // receive from the pipe

// Mutex — lock/unlock a shared resource
var mu sync.Mutex
mu.Lock()
  // only ONE goroutine can be here at a time
mu.Unlock()

// WaitGroup — wait for goroutines to finish
var wg sync.WaitGroup
wg.Add(1)       // "I'm adding 1 job"
go func() {
    defer wg.Done()  // "job is done"
    doWork()
}()
wg.Wait()        // "wait until all jobs are done"
```
