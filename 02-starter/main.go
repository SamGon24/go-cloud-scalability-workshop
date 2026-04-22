// 02-starter: YOUR TASK.
//
// Right now, jobs go INTO the queue but nothing processes them.
// The queue will fill up and new requests will block forever.
//
// Your job, in teams:
//   1. Add a worker goroutine that reads from jobQueue and "processes" each job.
//   2. Simulate processing time with time.Sleep(2 * time.Second).
//   3. Print something like "Worker processing job X" so you can see it work.
//   4. STRETCH: start multiple workers so jobs process in parallel.
//   5. STRETCH: give each job a unique ID so you can tell them apart.

package main

import (
	"fmt"
	"log"
	"net/http"
)

var jobQueue = make(chan string, 100)

// TODO: write a worker() function that loops over jobQueue and processes each job.

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	jobQueue <- "new job"
	fmt.Fprintf(w, "Job accepted!\n")
}

func main() {
	// TODO: start one or more workers here with the `go` keyword.

	http.HandleFunc("/upload", uploadHandler)
	log.Println("starter server listening on :8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
