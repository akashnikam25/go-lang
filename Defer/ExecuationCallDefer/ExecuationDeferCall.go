package main

import "fmt"

func main() {
	defer fmt.Println("First ")
	defer fmt.Println("Second ")
	defer fmt.Println("Third ")
	fmt.Println("Hi All")
	defer fmt.Println("Fourth")
	fmt.Println("Bye All")
}

// defer :- the defer keyword instructs a function to execute after the surrounding function completes.

// how multiple defer keyword works?->It works like stack
// Each time a function is deferred, it’s pushed onto a stack.
// Once the function execution completes, the deferred functions are popped,
// so that they execute in a “first-in-first-out” order:

//O/P :- For this Program
// Hi All
// Bye All
// Fourth
// Third
// Second
// First
