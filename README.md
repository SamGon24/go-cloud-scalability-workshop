# From Local Code → Scalable Cloud Service (with Go)

**Audience:** CS 3398 (Intro to Software Engineering)
**Duration:** 80 minutes

---

## Pre-Workshop Setup (for students)

```bash
# 1. Install Go 1.21+
# macOS:   brew install go
# Linux:   https://go.dev/dl/
# Windows: https://go.dev/dl/

# 2. Verify
go version   # should show go1.21+

# 3. Install a load-test helper (optional but useful)
# macOS: brew install hey
# or:    go install github.com/rakyll/hey@latest
```

No modules, no dependencies — every example is a single `main.go` you can run with:

```bash
go run main.go
```

---

## Folder Map

| Folder | Purpose | When to use |
|---|---|---|
| `01-naive/` | The blocking server (the "before" picture) | Live demo at 0:15 |
| `02-starter/` | Starter code students extend | Hands-on at 0:30 |
| `03-solution/` | Single-worker solution | Reveal at ~0:45 |
| `04-stretch/` | Multi-worker + graceful shutdown | Stretch goal |
| `05-bonus-kafka-style/` | Queue + producers + consumers pattern | "If we had Kafka" teaser |

Each folder has its own `README.md` with run instructions and teaching notes.

---

## Run of Show (condensed)

```
0:00  Kickoff                     (5 min)
0:05  Core concepts               (10 min)
0:15  Live demo: 01-naive         (15 min)
0:30  Hands-on: 02-starter        (25 min)
0:55  Architecture activity       (15 min)
1:10  Share + wrap                (10 min)
```
