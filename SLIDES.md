## From Local Code → Scalable Cloud Service

CS 3398 · Workshop · 80 minutes

*Subtitle:* "Today we'll take a simple app, break it, and redesign it like cloud engineers."

---

## What is scalability?

**Handling more users without falling over.**

Two kinds:

- **Vertical** — bigger box. Easy. Has a ceiling.
- **Horizontal** — more boxes. Hard. No ceiling.

---

## What breaks first?

> If your app gets 10,000 users tomorrow… what fails?

Likely culprits:

- CPU / memory on one box
- Database connections
- Head-of-line blocking on slow requests
- In-memory state that can't be shared

---

## Common mistakes

- Everything in memory on one server
- One instance, one process, one point of failure
- Blocking requests — user waits for slow work
- Tight coupling — API does the work itself

---

## What good looks like

- **Stateless services** — any copy can serve any request
- **Horizontal scaling** — add instances behind a load balancer
- **Async processing** — fast accept, slow process, separately
- **Separation of concerns** — API, workers, storage, queue

---

## Today's challenge

> You have a blocking Go server. One request = 2 seconds of work.

1. Watch it struggle (demo)
2. Turn it async (your task)
3. Design the "viral app" version (team activity)

---

## Live demo: the naive version

```go
func process(w http.ResponseWriter, r *http.Request) {
    time.Sleep(2 * time.Second)   // the enemy
    fmt.Fprintf(w, "Processed!\n")
}
```

---

## Your task (25 min)

**Turn blocking → async using a queue + workers.**

Starter code: `02-starter/main.go`

- Add a `worker()` goroutine that drains `jobQueue`
- Simulate 2s work with `time.Sleep`
- Start multiple workers from `main()`

Stretch: back-pressure (reject when full), job IDs.

---

## The pattern you just built

```
  HTTP handler → channel → worker pool
       (fast)   (queue)    (parallel)
```

In production, replace the channel with **Kafka / SQS / NATS / Redis Streams**.
The shape is identical.

---

## Architecture activity (15 min)

> "Your app goes viral. Users upload images that need ML processing. 10k uploads/hour. Design the system."

On the whiteboard:

1. Where does upload land?
2. Where does ML run?
3. What decouples them?
4. Where does data live?
5. What scales horizontally?

---

## Wrap

> "What you built today is the skeleton of every modern cloud system."
