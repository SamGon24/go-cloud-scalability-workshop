// 01-naive: The "before" picture.
//
// Every request blocks for 2 seconds. Run this, then hit it with 50
// concurrent requests (see ../load-test.sh). Watch latency balloon.
//
// This is the problem we will solve in 02-starter.

package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func process(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	time.Sleep(2 * time.Second) // simulate slow work (image resize, ML call, etc.)
	fmt.Fprintf(w, "Processed in %s\n", time.Since(start))
}

func main() {
	http.HandleFunc("/process", process)
	log.Println("naive server listening on :8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
