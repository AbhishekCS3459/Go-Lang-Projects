package main

import (
	"fmt"
	"time"
)

func main() {
	// input
	start_time:=time.Now()
	defer func(){
	 fmt.Println("Time taken:",time.Since(start_time))
	}()
	var input = []int{1, 2, 3, 4, 5}
	// stage 1: this stage will put this input to the channel
	dataChannel := sliceToChannel(input)
	// stage 2: this stage will square the data from the channel
	finnalChannel := sq(dataChannel)
	// stage 3: this stage will print the data from the channel
	printChannel(finnalChannel)
}

func sliceToChannel(input []int) chan int {
	ans := make(chan int)
	go func() {
		for _, value := range input {
			ans <- value
		}
		close(ans)
	}()
	return ans
}
func sq(in <- chan int) chan int{
	ans := make(chan int)
	go func() {
		for value := range in {
			ans <- value * value
		}
		close(ans)
	}()
	return ans
}

func printChannel(in <- chan int) {
	for value := range in {
		fmt.Println(value)
	}
}