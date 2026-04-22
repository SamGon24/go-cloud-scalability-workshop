# 01 — Naive Blocking Server (demo)

## Run it

```bash
cd 01-naive
go run main.go
```

## Test it (one request)

```bash
curl localhost:8081/process
# ~2 seconds later:
# Processed in 2.0012s
```

## Break it (fifty concurrent requests)

From a second terminal:

```bash
../load-test.sh
# OR, if you have `hey` installed:
hey -n 50 -c 50 http://localhost:8081/process
```

## What students should notice

- Go actually handles each request on its own goroutine, so all 50 finish in ~2s, not 100s. That's a nice surprise.
- BUT: the `time.Sleep` is standing in for real work (CPU, DB, disk, external API). In reality, 50 parallel 2s DB calls would saturate the connection pool and crash.
- The *client* still waits. The user experience is bad.
- We have no visibility, no retry, no back-pressure. One slow dependency takes the whole service down.

## Teaching prompts

> "Where is the state? (answer: nowhere yet — but watch what happens when you add an in-memory map)"
>
> "What if the request is a 100MB image upload — do we want the client to wait?"
>
> "How would you scale this across 3 machines? What breaks?"

## What comes next

`../02-starter/` — we accept the job immediately and process it in the background.
