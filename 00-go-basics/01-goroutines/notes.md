# Goroutines 🏃‍♂️

## What is a Goroutine?

Imagine you're a chef (the main program). Without goroutines, you do everything yourself:
1. Chop vegetables ✂️
2. Wait...
3. Boil water 🫧
4. Wait...
5. Fry meat 🥩

This is **slow** because you wait for each task before starting the next.

With goroutines, you **hire more chefs** (goroutines) and they all work at the same time:
```
Chef 1 (main)    → manages the kitchen
Chef 2 (goroutine) → chopping vegetables RIGHT NOW
Chef 3 (goroutine) → boiling water RIGHT NOW
Chef 4 (goroutine) → frying meat RIGHT NOW
```

All happening **simultaneously**. Much faster!

---

## The Syntax

```go
// Normal function call — you WAIT for it to finish
doWork()

// Goroutine — it runs in BACKGROUND, you don't wait
go doWork()
```

Just add the word `go` in front. That's it!

---

## Your First Goroutine

```go
package main

import (
    "fmt"
    "time"
)

func sayHello(name string) {
    fmt.Println("Hello,", name)
}

func main() {
    go sayHello("Alice")  // runs in background
    go sayHello("Bob")    // runs in background
    go sayHello("Charlie") // runs in background

    // ⚠️ PROBLEM: main() might finish before goroutines run!
    time.Sleep(1 * time.Second)  // wait a bit so goroutines can finish
    fmt.Println("Done!")
}
```

**Output (order may vary!):**
```
Hello, Bob
Hello, Alice
Hello, Charlie
Done!
```

The order is **random** because all three run at the same time and we don't know who finishes first.

---

## The Problem: Race to the Exit

```go
func main() {
    go sayHello("Alice")
    // main() ends here — Alice never gets to say hello! 😱
}
```

When `main()` exits, ALL goroutines are killed immediately, even if they haven't finished.

**Solutions:**
1. `time.Sleep()` — hacky, don't use in real code
2. `sync.WaitGroup` — the RIGHT way (covered in 03-mutex section)
3. `channels` — another great way (covered in 02-channels section)

---

## Why Goroutines are Awesome

Normal OS threads use ~1MB of memory each. If you need 10,000 threads, that's 10GB of RAM! 😱

Goroutines use only ~2KB each. You can have **millions** of goroutines without breaking a sweat. This is why Go is perfect for distributed systems where you handle thousands of connections at once.

---

## Common Mistake: Goroutine with a Loop

```go
// ❌ WRONG — all goroutines will print the SAME value (the final i)
for i := 0; i < 3; i++ {
    go func() {
        fmt.Println(i)  // captures i by reference!
    }()
}
// Prints: 3, 3, 3  (probably)

// ✅ CORRECT — pass i as a parameter
for i := 0; i < 3; i++ {
    go func(num int) {
        fmt.Println(num)  // has its own copy of i
    }(i)
}
// Prints: 0, 1, 2  (in some order)
```

This is called a **closure bug** — super common mistake, remember it!

---

## Summary

| Concept | Key Point |
|---------|-----------|
| `go func()` | Runs a function in the background |
| Order | Goroutines run in **unpredictable order** |
| Lifecycle | Goroutines die when `main()` exits |
| Cost | Very cheap — you can have millions |
| Closure bug | Always pass loop variables as parameters |

---

## Next Step

Now try the exercise! → [exercise.go](./exercise.go)
After trying, check → [solution.go](./solution.go)
