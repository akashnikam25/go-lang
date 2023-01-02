package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	http.ListenAndServe(":8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		fmt.Fprintf(os.Stdout, "Request is Processing\n")
		select {
		case <-time.After(2 * time.Second):
			w.Write([]byte("Reuest Processed"))
		case <-ctx.Done():
			fmt.Fprintf(os.Stderr, "Request cancelled \n")
		}
	}))
}
