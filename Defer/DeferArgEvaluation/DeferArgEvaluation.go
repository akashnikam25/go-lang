package main

import "fmt"

var i int

func increment() int {
	i++
	return i
}
func main() {
	defer fmt.Println("i defer	:	", increment())
	fmt.Println("i count :	", increment())
}
//If you want to evaluate everything when the deferred function executes, 
//you should ensure all functions are called within the deferred function call.