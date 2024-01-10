// select statement in Go
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("select statement in Go")

	myChannel := make(chan string)
	secondChannel := make(chan string)

	go func() {
		myChannel <- "hello"
	}()
	go func() {
		secondChannel <- "world"
	}()
	select {
	case msgfromMychannel := <-myChannel:
		fmt.Println(msgfromMychannel)
	case msgfromAnother := <-secondChannel:
		fmt.Println(msgfromAnother)
	}

	time.Sleep(time.Second * 2)
}
