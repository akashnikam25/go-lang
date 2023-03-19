package main

import "fmt"

// func main1() {
// 	n := 3
// 	out := make(chan int)

// 	go getmultiplication(n, out)
// 	fmt.Println("multiplication", <-out)
// }

// func getmultiplication(n int, out chan int) {
// 	res := n * n
// 	out <- res
// }

func main() {
	out := make(chan int)
	in := make(chan int)

	// Create 3 `multiplyByTwo` goroutines.
	go multiplyByTwo(in, out)
	go multiplyByTwo(in, out)
	go multiplyByTwo(in, out)

	// Up till this point, none of the created goroutines actually do
	// anything, since they are all waiting for the `in` channel to
	// receive some data, we can send this in another goroutine
	go func() {
		in <- 1
		in <- 2
		in <- 3
		in <- 4
	}()

	// Now we wait for each result to come in
	fmt.Println(<-out)
	fmt.Println(<-out)
	fmt.Println(<-out)
	fmt.Println(<-out)
}

func multiplyByTwo(in <-chan int, out chan<- int) {
	fmt.Println("Initializing goroutine...")
	for {
		num := <-in
		result := num * 2
		out <- result
	}
}
