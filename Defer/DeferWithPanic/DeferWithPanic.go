package main

import "fmt"

func main() {
	defer fmt.Println("after panic happen")

	fmt.Println("hello World")

	panic("Somthing went Wrong")
}
/*
Output 
hello World
after panic happen
panic: Somthing went Wrong

goroutine 1 [running]:
main.main()
	d:/Akash/Go/Basic-Go-lang/Defer/DeferWithPanic/DeferWithPanic.go:10 +0xb8
exit status 2*/

//to avoid abonarmally termination from panic we can use recover method 

func main1() {
	defer func() {
		// we can use the recover function inside the deferred function
  		// to tell if it was called after `panic` was called
		if r := recover(); r != nil {
			fmt.Println("Panic Recoverd	:	", r)
		}
	}()

	fmt.Println("hello World")

	panic("Somthing went Wrong")
}

