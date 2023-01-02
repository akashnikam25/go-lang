package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func operation1(ctx context.Context) (err error) {
	fmt.Println("Opearation1 started")
	//time.Sleep(5 * time.Second)
	fmt.Println("Ctx key Value :	", ctx.Value("1"))
	err = errors.New("error")
	//err = nil
	return
}

// operation3 func will print ctx value
func operation3(ctx context.Context) {
	fmt.Println("Ctx key Value operation3	:	", ctx.Value("1"))
}
func operation2(ctx context.Context) {
	fmt.Println("operation2 startrd		")
	select {
	case <-time.After(500 * time.Millisecond):
		fmt.Println("Operation2 completed")
		operation3(ctx)
	case <-ctx.Done():
		fmt.Println("all Opeartion Cancelled ")
	}
}
func main() {
	ctx := context.Background()
	ctxWithValue := context.WithValue(ctx, "1", "Akash")
	ctxCancel, ctxCancelFn := context.WithCancel(ctxWithValue)

	go func() {
		err := operation1(ctxCancel)
		//if operation1 returns the error then cancel all opeartions
		if err != nil {
			fmt.Println("err	:		", err)
			ctxCancelFn()
		}
	}()
	//run operation2
	operation2(ctxCancel)

}
