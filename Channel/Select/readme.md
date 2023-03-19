# Channel Select Example in Go
This Go program demonstrates the use of channels and select statement to concurrently execute two functions and receive their results.

The program defines two functions, fast and slow, which both take an integer and a channel as input. These functions perform the same operation of doubling the input integer, but slow takes longer to complete due to a longer time.Sleep duration.

In the main function, two channels are created to receive the results of the fast and slow functions. Both functions are then started in separate goroutines using the go keyword, allowing them to execute concurrently.

A select statement is used to receive the first result that becomes available on either channel. Depending on which channel receives a result first, the program prints a message indicating whether fast or slow finished first, along with the result.

This program demonstrates how channels and the select statement can be used to efficiently execute multiple functions concurrently and receive their results as soon as they are available.

## Usage
To run the program, save the code to a file named main.go, and then run the following command:
```
go run main.go
```
The program should output a message indicating which function finished first, along with the result.