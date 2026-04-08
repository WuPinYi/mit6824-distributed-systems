# Channels 📬

## What is a Channel?

A channel is like a **mailbox** between goroutines.

```
Goroutine A  →→→  [📬 channel]  →→→  Goroutine B
   (sender)                          (receiver)
```

- Goroutine A puts a message in the mailbox
- Goroutine B picks up the message from the mailbox
- They never need to share memory directly — much safer!

---

## The Syntax

```go
// Create a channel that carries integers
ch := make(chan int)

// Send the number 42 into the channel
ch <- 42

// Receive a value from the channel
value := <-ch
fmt.Println(value) // 42
```

The arrow `<-` shows the direction of data flow.

---

## Channels BLOCK (this is important!)

By default, channels are **synchronous**:

- **Sending blocks** until someone is ready to receive
- **Receiving blocks** until someone is ready to send

Think of it like a phone call — both people need to be there at the same time.

```go
ch := make(chan int)

// This would DEADLOCK if you run it in main alone:
ch <- 42  // blocks forever, nobody is receiving!

// The right way — sender and receiver in different goroutines:
go func() {
    ch <- 42  // goroutine sends, then waits for receiver
}()
value := <-ch  // main receives, unblocks the goroutine
fmt.Println(value) // 42
```

---

## Buffered Channels

What if you don't want to block immediately? Use a **buffered channel** with a capacity:

```go
// Buffered channel — can hold up to 3 values without blocking
ch := make(chan int, 3)

ch <- 1  // doesn't block (buffer has space)
ch <- 2  // doesn't block (buffer has space)
ch <- 3  // doesn't block (buffer has space)
ch <- 4  // BLOCKS! buffer is full

v := <-ch  // receive 1, now there's space
ch <- 4    // now this works
```

Think of buffered channels as a **queue with a size limit**.

---

## Closing a Channel

When you're done sending, close the channel:

```go
ch := make(chan int, 5)
ch <- 10
ch <- 20
ch <- 30
close(ch)  // signal: no more values coming

// Loop over all values until channel is closed
for v := range ch {
    fmt.Println(v)  // prints 10, 20, 30
}
```

---

## Using Channels to Replace time.Sleep

Remember in goroutines we used `time.Sleep` to wait? Channels are the better way:

```go
// ❌ Hacky way
go doWork()
time.Sleep(1 * time.Second) // hope it finishes in time...

// ✅ Channel way — guaranteed to wait for the result
done := make(chan bool)
go func() {
    doWork()
    done <- true  // signal: I'm done!
}()
<-done  // blocks until doWork() finishes
fmt.Println("Work is done!")
```

---

## Real Example: Parallel Web Requests

Imagine fetching data from 3 servers at once:

```go
func fetchData(serverName string, ch chan string) {
    // simulate fetching data
    time.Sleep(100 * time.Millisecond)
    ch <- serverName + ": data ready!"
}

func main() {
    ch := make(chan string, 3)  // buffered for 3 results

    // Launch 3 fetches at the same time
    go fetchData("Server A", ch)
    go fetchData("Server B", ch)
    go fetchData("Server C", ch)

    // Collect all 3 results
    for i := 0; i < 3; i++ {
        result := <-ch
        fmt.Println(result)
    }
}
```

All 3 fetches happen in parallel — instead of 300ms total, it takes only ~100ms! 🚀

---

## Select Statement

What if you have multiple channels and want to handle whichever one is ready first?

```go
ch1 := make(chan string)
ch2 := make(chan string)

go func() { ch1 <- "from channel 1" }()
go func() { ch2 <- "from channel 2" }()

// select picks whichever channel is ready first
select {
case msg1 := <-ch1:
    fmt.Println("Received:", msg1)
case msg2 := <-ch2:
    fmt.Println("Received:", msg2)
}
```

`select` is like a `switch` statement but for channels. You'll use this **a lot** in Raft (Lab 2)!

---

## Summary

| Concept | Key Point |
|---------|-----------|
| `make(chan T)` | Create an unbuffered channel of type T |
| `make(chan T, n)` | Create a buffered channel (capacity n) |
| `ch <- value` | Send into channel (may block) |
| `value := <-ch` | Receive from channel (may block) |
| `close(ch)` | Signal no more values will be sent |
| `for v := range ch` | Loop until channel is closed |
| `select` | Handle whichever channel is ready first |

---

## Next Step

Try the exercise! → [exercise.go](./exercise.go)
After trying, check → [solution.go](./solution.go)
