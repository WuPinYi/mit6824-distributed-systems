# Mutex & WaitGroup 🔐

## The Problem: Race Conditions

Imagine two cashiers at a bank. Customer A has $100 in their account.
Both cashiers process a $80 withdrawal at the SAME TIME:

```
Cashier 1: reads balance → $100
Cashier 2: reads balance → $100
Cashier 1: subtracts $80 → $20, writes back
Cashier 2: subtracts $80 → $20, writes back  ← WRONG! Should be -$60!
```

The final balance is $20 instead of -$60. The bank just lost $80! 💸

This is a **race condition** — two goroutines reading and writing the same data at the same time causes bugs.

---

## What is a Mutex?

A mutex (mutual exclusion) is like a **bathroom key** 🔑

```
[Bathroom = shared data]

Person 1 (goroutine 1): takes key, enters bathroom, does work, leaves, returns key
Person 2 (goroutine 2): waits for the key... gets key, enters, does work, leaves
```

Only ONE person can be in the bathroom at a time. Everyone else waits.

---

## The Syntax

```go
import "sync"

var mu sync.Mutex  // declare a mutex

// Safe way to modify shared data:
mu.Lock()          // grab the key — if someone else has it, WAIT
  balance -= 80    // do the work safely
mu.Unlock()        // return the key — let others in

// Even better — use defer so you never forget to Unlock:
mu.Lock()
defer mu.Unlock()  // will run at end of function, even if panic occurs
  balance -= 80
```

---

## Race Condition Example

```go
// ❌ BAD: race condition
counter := 0
for i := 0; i < 1000; i++ {
    go func() {
        counter++  // NOT SAFE! Multiple goroutines write at same time
    }()
}
time.Sleep(time.Second)
fmt.Println(counter)  // Prints something < 1000! Data was lost.

// ✅ GOOD: protected with mutex
var mu sync.Mutex
counter := 0
for i := 0; i < 1000; i++ {
    go func() {
        mu.Lock()
        counter++  // only one goroutine can be here at a time
        mu.Unlock()
    }()
}
time.Sleep(time.Second)
fmt.Println(counter)  // Always prints 1000 ✓
```

---

## WaitGroup: Waiting for Goroutines Properly

We've been using `time.Sleep()` to wait for goroutines — that's unreliable.
`sync.WaitGroup` is the proper solution:

```go
var wg sync.WaitGroup

// For each goroutine you launch, call wg.Add(1)
wg.Add(1)
go func() {
    defer wg.Done()  // call Done when goroutine finishes (decrements counter)
    doWork()
}()

wg.Add(1)
go func() {
    defer wg.Done()
    doMoreWork()
}()

wg.Wait()  // blocks here until counter reaches 0
fmt.Println("All goroutines finished!")
```

Think of it like a **scoreboard**:
- `wg.Add(1)` → "one more job to wait for"
- `wg.Done()` → "one job finished" (like subtracting 1)
- `wg.Wait()` → "wait until score hits 0"

---

## Combining Mutex + WaitGroup (Very Common Pattern)

```go
var (
    mu      sync.Mutex
    wg      sync.WaitGroup
    results []int
)

for i := 0; i < 5; i++ {
    wg.Add(1)
    go func(n int) {
        defer wg.Done()

        result := n * n  // compute (safe, each goroutine has its own n)

        mu.Lock()
        results = append(results, result)  // writing to shared slice — NEEDS LOCK
        mu.Unlock()
    }(i)
}

wg.Wait()
fmt.Println(results)  // [0 1 4 9 16] in some order
```

---

## RWMutex: Reads are Free (Almost)

What if many goroutines are **reading** and only occasionally **writing**?
A regular Mutex makes every reader wait. That's slow.

`sync.RWMutex` is smarter:
- **Multiple readers** can hold the lock at the same time (no waiting!)
- **Only ONE writer** can hold the lock (and blocks all readers)

```go
var rwmu sync.RWMutex
data := make(map[string]int)

// Reading — use RLock (many goroutines can do this simultaneously)
rwmu.RLock()
value := data["key"]
rwmu.RUnlock()

// Writing — use Lock (exclusive, blocks everyone)
rwmu.Lock()
data["key"] = 42
rwmu.Unlock()
```

You'll use `RWMutex` in the Raft labs where many things read state but only the leader writes.

---

## Detecting Race Conditions

Go has a built-in **race detector**! Run your code like this:

```bash
go run -race yourfile.go
```

It will warn you if two goroutines access the same variable without a lock. Use this constantly when writing lab code!

---

## Summary

| Concept | What it does | When to use |
|---------|-------------|-------------|
| `sync.Mutex` | Only 1 goroutine can be in the critical section | Protecting shared data |
| `mu.Lock()` / `mu.Unlock()` | Grab / release the lock | Around code that touches shared data |
| `defer mu.Unlock()` | Auto-unlock when function exits | Always prefer this! |
| `sync.WaitGroup` | Wait for N goroutines to finish | Replace `time.Sleep` |
| `wg.Add(n)` | Expect n more goroutines | Before launching goroutines |
| `wg.Done()` | This goroutine is done | At the end of each goroutine |
| `wg.Wait()` | Block until all are done | After launching all goroutines |
| `sync.RWMutex` | Many readers OR one writer | Read-heavy workloads |

---

## Next Step

Try the exercise! → [exercise.go](./exercise.go)
After trying, check → [solution.go](./solution.go)

Then you're ready for **Lab 1: MapReduce**! 🎉
