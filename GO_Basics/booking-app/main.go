package main

import (
	"fmt" // Package fmt implements formatted I/O with functions analogous to C's printf and scanf.
	"sync"
	"time"
)


type UserType struct {
	name            string
	age             uint
	email           string
	numberofTickets uint
	isUserOpted     bool
}

var wg = sync.WaitGroup{}

func main() {
	var Data = make([]UserType, 0)
	var userData = UserType{
		name:            "Jon",
		age:             13,
		email:           "xyz@gmail.com",
		numberofTickets: 1,
		isUserOpted:     true,
	}
	Data = append(Data, userData)
	fmt.Println(Data)
	fmt.Println("we are sending 0th tickets you dont have to wait for it")
	// wg.Add(1)
	// go sendTicket(&Data[0])
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go sendTicket(&Data[0])
	}
	fmt.Println("lets continue some other work while its done")

		wg.Wait()
}

func sendTicket(Data *UserType) {
	var ticketInfo = fmt.Sprintf(" %v Ticket for %v having email: %v  ", Data.numberofTickets, Data.name, Data.email)
	var idx int64 = 0
	fmt.Println("sending ticket...")
	 start_Time:= time.Now()
	for idx < 1e10 {
		idx++
	}
	fmt.Println("time taken to send ticket is ", time.Since(start_Time).Seconds(), "sec")
	fmt.Println(ticketInfo)
	wg.Done()
}
