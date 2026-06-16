# Day 1 - Goroutines and Go Scheduler

## Objective

Learn the basics of Go concurrency using Goroutines and understand how the Go Scheduler executes multiple tasks concurrently.

## Programs Included

### 1. Basic Goroutine Example

* Creates a single goroutine.
* Demonstrates concurrent execution alongside the main function.

### 2. Multiple Goroutines

* Launches multiple goroutines simultaneously.
* Shows how tasks run independently and concurrently.

### 3. Goroutines with WaitGroup

* Uses `sync.WaitGroup` to synchronize goroutines.
* Ensures the main function waits for all goroutines to complete.

### 4. Goroutines with Execution Time Measurement

* Runs multiple workers concurrently.
* Measures total execution time using the `time` package.

## Prerequisites

* Go 1.18 or later
* VS Code or any Go-supported IDE
* Basic knowledge of Go syntax

## How to Run

1. Open a terminal in the project directory.

2. Verify Go installation:

```bash
go version
```

3. Run a program:

```bash
go run filename.go
```

Example:

```bash
go run goroutine_waitgroup.go
```

## Key Concepts Learned

### Goroutines

A goroutine is a lightweight thread managed by the Go runtime.

Example:

```go
go myFunction()
```

### WaitGroup

WaitGroup is used to wait for multiple goroutines to finish execution.

Example:

```go
var wg sync.WaitGroup
wg.Add(1)
go worker(&wg)
wg.Wait()
```

### Go Scheduler

The Go Scheduler manages goroutines and efficiently maps them to operating system threads.

## Expected Learning Outcomes

After completing these programs,I was able to:

* Create and execute goroutines.
* Run multiple tasks concurrently.
* Synchronize goroutines using WaitGroup.
* Understand the role of the Go Scheduler.
* Measure execution time of concurrent programs.

