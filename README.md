# Goroutine Mastery – Practical Examples for Deep Understanding

This repository contains practical and unique Go code examples to help you master goroutines from basics to advanced internals.  
Based on the article: [Inside Goroutines: From Zero to Senior-Level Mastery](https://medium.com/@sogol.hedayatmanesh/inside-goroutines-from-zero-to-senior-level-mastery-ed0f3a465abf)

---

## Project Structure

### 1. basics/
Basic examples introducing goroutines and their behavior.

- `hello.go`  
  Simple “Hello from goroutine” example demonstrating concurrent execution.  
  **Run:**
  ```bash
  go run basics/hello.go
  ```
- `race-condition.go`
  Shows a common data race issue when multiple goroutines access shared memory without synchronization.

  **Run:**
  ```bash
  go run basics/race-condition.go
  ```
### 2. lifecycle/
   Examples demonstrating goroutine lifecycle states through a worker pool.

- `worker_pool.go`
  Multiple worker goroutines process jobs from a channel, illustrating the New, Runnable, Running, Blocked, and Dead states.

  **Run:**
  ```bash
  go run lifecycle/worker_pool.go
  ```
### 3. stack_management/
   Exploring goroutine stack growth and dynamic memory management.
- `stack_growth.go`
  Uses deep recursion to force stack growth in a goroutine.
  **Run:**
  ```bash
  go run stack_management/stack_growth.go
  ```
### 4. benchmarks/
   Performance comparison focusing on goroutine creation and concurrency.
- `goroutine_vs_thread.go`
   Creates 100,000 goroutines to benchmark lightweight concurrency in Go.
  **Run:**
  ```bash
  go run benchmarks/goroutine_vs_thread.go
  ```

## How to Use
  Clone the repository:
  ```bash
      git clone https://github.com/sogolHedaytaManesh/go-goroutine-mastery.git
      cd go-goroutine-mastery
  ```
Run the examples with go run as shown above.

Experiment with code and modify examples to deepen your understanding.

## Why This Repository?
Goroutines are at the core of Go’s concurrency model, but most resources only scratch the surface. This repo provides hands-on examples aligned with the detailed article to give you real insight into goroutine internals, lifecycle, stack management, and performance.

##Learn More
Read the full article here:
[Inside Goroutines: From Zero to Senior-Level Mastery
](https://medium.com/@sogol.hedayatmanesh/inside-goroutines-from-zero-to-senior-level-mastery-ed0f3a465abf)
