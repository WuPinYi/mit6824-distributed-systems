# MIT 6.824 Distributed Systems — Learning Journey 🚀

---

## What is this course?

MIT 6.824 teaches you how to build systems that run across **many computers at once** — and keep working even when some of those computers crash. This is how real companies like Google, Amazon, and Netflix build their infrastructure.

You'll build real systems in **Go (Golang)**.

---

## The Big Picture

```
Single Computer  →  Many Computers Working Together
     💻               💻 💻 💻 💻 💻
                   (even if some crash, it still works!)
```

The three hardest problems in distributed systems:
1. **Crashes** — a server can die at any moment
2. **Network problems** — messages get lost or delayed
3. **Agreement** — how do all servers agree on the same truth?

This course teaches you to solve all three.

---

## Roadmap

| Stage | Topic | Status |
|-------|-------|--------|
| 🐣 00 | Go Basics (goroutines, channels, mutex) | 🔄 In Progress |
| 🗺️ Lab 1 | MapReduce | ⏳ Pending |
| 🗳️ Lab 2 | Raft Consensus (Leader Election + Log Replication) | ⏳ Pending |
| 🗄️ Lab 3 | Fault-Tolerant Key/Value Store | ⏳ Pending |
| 🌐 Lab 4 | Sharded Key/Value Store | ⏳ Pending |

---

## Folder Structure

```
mit6824/
├── 00-go-basics/          ← Start here! Go skills for the labs
│   ├── 01-goroutines/     ← Lightweight threads
│   ├── 02-channels/       ← How goroutines talk to each other
│   └── 03-mutex/          ← Protecting shared data
```

---

## How to Use This Repo

Each folder has:
- `notes.md` — Super simple explanation
- `exercise.go` — Try to write this yourself first!

---

## Setup

```bash
# Check Go is installed
go version  # should show go1.21 or higher

# Run any exercise
go run exercise.go

```

---

## Resources

- [MIT 6.824 Course Page](https://pdos.csail.mit.edu/6.824/)
- [Go Tour (official)](https://go.dev/tour/)
- [Effective Go](https://go.dev/doc/effective_go)

---

*"The goal is not to memorize. The goal is to understand WHY."* 🧠
