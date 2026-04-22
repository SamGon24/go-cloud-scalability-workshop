# Instructor Overview

Everything you need to run this workshop without looking anything up. Use this in presenter mode.

---

## Timing

```
0:00  0:05   Kickoff                       -
0:05  0:15   Concepts (slides 1-4)         -
0:15  0:30   Demo 01-naive                 ← run load-test.sh
0:30  0:55   Hands-on 02-starter           ← students code
              0:30 hand out + explain
              0:35-0:50 teams work
              0:50 reveal 03-solution
              0:55 explain back-pressure
0:55  1:10   Architecture activity         ← whiteboard
1:10  1:20   Share-back + wrap             -
```

To save on time if needed: skip 04-stretch walkthrough or cut architecture activity discussion down.

---

## Section-by-section answer key

### 0:05 — "What breaks first at 10k users?"

**Expected student answers + what to validate:**

| Student says | You say |
|---|---|
| "The server crashes" | Good — which resource runs out first? (CPU? memory? file handles? DB connections?) |
| "Requests slow down" | Right — and why? (head-of-line blocking, one slow request blocks the goroutine/thread) |
| "Database" | Best answer — because a stateless app scales horizontally, but a DB doesn't, easily |
| "Network" | True but usually last — dismiss gently |
| "Memory" | Only if they're storing things per-user in memory — good lead-in to "stateless" |

### 0:15 — Live demo of `01-naive/`

Commands, in order:

```bash
cd 01-naive
go run main.go
```

Second terminal:

```bash
curl localhost:8081/process      # one request, 2s
../load-test.sh                  # fifty, watch the wall-clock
```

**What you're really showing them:**

- Go is unusually forgiving here — all 50 run in parallel goroutines, so wall-clock is ~2s.
- **Don't let that trick them.** Say:
  > "If the 2-second wait was a real database query, you'd blow through your connection pool at request 21. The `time.Sleep` is a lie — real work competes for real resources."
- The *client* still waits 2s for every request. That's bad UX even when the server holds up.

### 0:30 — Hand out `02-starter/`

Read this aloud:

> "I want you to turn the blocking version into an async one. When someone uploads, the server says 'got it!' instantly. In the background, workers do the real work. You have 20 minutes. Peek at `03-solution` only when you're stuck."

### 0:50 — Reveal `03-solution/main.go`

Walk through **in this order** (most valuable first):

1. **The `select` with `default`** — this is back-pressure.
   > "This one pattern is worth the whole workshop. When a real system gets overloaded, it must choose: reject fast, or fall over. This picks the former."
2. **`atomic.Uint64`** — counters without mutexes.
3. **Worker pool sized to 3** — not infinite. Why? Downstream resources.

### 0:55 — Architecture activity

See [ARCHITECTURE.md](./ARCHITECTURE.md) for the prompt and the expected diagram.

---
