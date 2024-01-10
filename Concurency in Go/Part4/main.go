package main

import (
	"fmt"
	"math/rand"
)

func main() {
	channel := make(chan string)
	
	go throwingNinjaStar(channel)

	// for message:=range channel {
	// 	fmt.Println(message)
	// }
//   or
	for{
		message,open:=<-channel
		if !open{
			break
		}
		fmt.Println(message)
	}
}
func throwingNinjaStar(channel chan string) {
	rounds := 30
	for idx := 0; idx < rounds; idx++ {
		channel <- fmt.Sprint("Ninja star attacked ", rand.Intn(10), " evil ninjas!")
	}
   close(channel)
}
