package main

import "fmt"

func main() {
	var username string = "Abhishek"
	fmt.Printf("Varianble is of type %T\n", username)
	var isLoggedIn bool = true
	fmt.Printf("Varianble is of type %T\n", isLoggedIn)
	var smallVal uint8 = 255 // 
	fmt.Print(smallVal)
   var smallFloat float64=23.23233333333333
   fmt.Print(smallFloat)


   // default values and some aliases

   var defaultvalue int
   fmt.Println(defaultvalue)
   fmt.Printf("variable is of type : %T \n",defaultvalue)
   
}
