#  Multiple Goroutines Example in Go
This Go program demonstrates the use of multiple goroutines to concurrently execute a function across multiple inputs.

The program defines a function multiplyByTwo which takes an input channel and an output channel as arguments. This function reads from the input channel, multiplies the received number by two, and sends the result to the output channel.

In the main function, three instances of multiplyByTwo are started in separate goroutines, each with their own output channel. A separate goroutine is created to send input values to the input channel.

Once all the goroutines are started, the program waits for each output channel to receive a result using the <-out operator. This causes the program to block until the result is received.

The multiplyByTwo function is designed to run indefinitely until the input channel is closed. This allows it to keep processing input values even after all input values have been sent.

This program demonstrates how goroutines can be used to process input values concurrently and efficiently, improving the overall performance of the program.

Usage
To run the program, save the code to a file named main.go, and then run the following command:

```
go run main.go
```
The program should output the result of multiplyByTwo function for each input value sent to the input channel.