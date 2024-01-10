package main

import (
	"fmt"
	"sync"
	"time"
)

var ws = sync.WaitGroup{}

func main() {
	start_time := time.Now()
	defer func() {
		fmt.Println("Time taken:", time.Since(start_time))
	}()
	evilNinjas := []string{"Tommy", "Johnny", "Bobby", "Billy"}
	evilSignal := make(chan bool)
	for _, ninja := range evilNinjas {
		ws.Add(1)
		go attack(ninja, evilSignal)
		fmt.Println(ninja, <-evilSignal)
	}
	ws.Wait()
}

func attack(target string, attacked chan bool) {
	fmt.Println("Attacked by", target)
	time.Sleep(time.Second)
	attacked <- true
	ws.Done()
}
