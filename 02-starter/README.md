# 02 — Starter (YOUR TASK)

You have **25 minutes** in your team.

## Goal

Turn the blocking server from `01-naive/` into an **async** system:
the client gets an immediate "Job accepted!" response, and workers
process the jobs in the background.

## Run it (before changes)

```bash
cd 02-starter
go run main.go
```

```bash
# in another terminal:
curl -X POST localhost:8081/upload
# Response: "Job accepted!"  (instant!)
```

But look at the server log — **nothing is processing jobs**. The queue is filling up.

## Your task

1. Write a `worker()` function that reads jobs from `jobQueue` and "processes" them.
2. Simulate processing with `time.Sleep(2 * time.Second)`.
3. Print when a worker starts and finishes a job.
4. Start the worker from `main()` with `go worker()`.

### Stretch goals

- Start **3** workers — all reading from the same channel. Watch them balance load.
- Give each job a unique ID (hint: `fmt.Sprintf("job-%d", counter)`).
- Give each worker an ID so you can see which one picked up which job.

## Test your work

```bash
# fire 10 requests
for i in $(seq 1 10); do curl -X POST localhost:8081/upload & done
wait
```

You should see the server accept all 10 instantly, then workers chew through them.

## Stuck?

Hints:

```go
// worker signature
func worker(id int) {
    for job := range jobQueue {
        // do the work
    }
}

// starting workers in main()
for i := 1; i <= 3; i++ {
    go worker(i)
}
```

Peek at `../03-solution/` only after you've tried.
