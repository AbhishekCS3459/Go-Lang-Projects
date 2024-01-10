package main

import (
	// users "booking-app/Users"
	// "booking-app/displayuser"
	"fmt" // Package fmt implements formatted I/O with functions analogous to C's printf and scanf.
	"sync"
	// "strconv"
	"time"
)

// func inputCandidateNames(CandidateNames *[]string, n int) {
// 	fmt.Println("Enter the names of the participants:")
// 	for i := 0; i < n; i++ {
// 		var name string
// 		fmt.Scan(&name)
// 		*CandidateNames = append(*CandidateNames, name)
// 	}
// }

// func greetUser(name *string, CandidateNames *[]string) {
// 	fmt.Println("Welcome inside function", *name)
// 	*name = "shyam"
// 	fmt.Print("the name of the candidates are ", *CandidateNames)
// }

type UserType struct {
	name            string
	age             uint
	email           string
	numberofTickets uint
	isUserOpted     bool
}

var wg = sync.WaitGroup{}

func main() {
	//		conferenceName := "Jon doe"
	//		const tickets int = 100
	//		fmt.Println("Welcome to ", conferenceName, "booking application!")
	//		fmt.Println("Available tickets are ", tickets)
	//		fmt.Printf("Available tickets are %d\n", tickets)
	//		var UserName string
	//		fmt.Println("Welcome ", UserName)
	//		fmt.Printf("type of the ticket is %T\n", tickets)
	//		fmt.Scan(&UserName)
	//		fmt.Println("Welcome ", UserName)
	//	   fmt.Print(&conferenceName)

	// var firstName, LastName string
	// var age uint
	// fmt.Println("Enter your first name")
	// fmt.Scan(&firstName)
	// fmt.Println("Enter your last name")
	// fmt.Scan(&LastName)
	// fmt.Println("Enter your age")
	// fmt.Scan(&age)
	// fmt.Println("Welcome ", firstName, LastName, "your age is ", age)

	// arrays and slices
	// var names = [5]string{"Jon", "Doe", "John", "Doe", "Jane"}
	// fmt.Println(names)
	// for i := 0; i < len(names); i++ {
	// 	fmt.Print(names[i], " ")
	// }
	// // slice in go
	// var namesSlice = []string{} // or nameSlice := []string{}
	// fmt.Printf("Enter the names of the participants:")
	// for i := 0; i < 5; i++ {
	// 	var name string
	// 	fmt.Scan(&name)
	// 	namesSlice = append(namesSlice, name)
	// }
	// // for each loop in go
	// for _, name := range namesSlice {
	// 	fmt.Println(strings.Fields(name))
	// }

	// // how for loop is working like a while loop
	// idx := 2
	// for idx >= 0 {
	// 	var name string
	// 	fmt.Println("Enter your name")
	// 	fmt.Scan(&name)
	// 	var age uint
	// 	fmt.Println("Enter your age")
	// 	fmt.Scan(&age)
	// 	fmt.Println("Welcome ", name, "your age is ", age)
	// 	idx--
	// }

	// // if else in go
	// var name string
	// fmt.Println("Enter your name")
	// fmt.Scan(&name)
	// if strings.Compare(name, "Jon") == 0 {
	// 	fmt.Println("Welcome ", name)
	// } else if strings.Compare(name, "Doe") == 0 {
	// 	fmt.Println("Welcome ", name)
	// } else {
	// 	fmt.Println("You are not welcome")
	// }

	// // switch case in go
	// switch name {
	// case "Jon":
	// 	fmt.Println("Welcome ", name)
	// case "Doe":
	// 	fmt.Println("Welcome ", name)
	// default:
	// 	fmt.Println("You are not welcome")
	// }
	// inputCandidateNames(&namesSlice, 2)
	// greetUser(&name, &namesSlice)
	// displayuser.PrintUser(&name)
	// users.ChangeUser(&name, "shyam")
	// displayuser.PrintUser(&name)

	// create a map for users
	// var userData=make(map[string]string)
	// userData["name"]="Jon"
	// userData[strconv.FormatUint(uint64(13),10)]="dare";
	// fmt.Printf(userData["13"])

	// create a list of maps
	// var usersData2=make([]map[string]string,10)
	// usersData2[0]=make(map[string]string)
	// usersData2[0]["name"]="Jon"
	// usersData2[0]["age"]="13"
	// fmt.Printf(usersData2[0]["age"])

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
