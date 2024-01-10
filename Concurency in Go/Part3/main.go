package main

import "fmt"

func main() {
	channels := make(chan string,20)
	// this is seperate go routine which do one thing putting this message to channel and then end itself
	//  go func() {
		idx:=10
		for idx >0 {
			channels <- "hello"
			idx--
		}
	// }()

	 
	fmt.Println("recived form channel",<-channels)

}